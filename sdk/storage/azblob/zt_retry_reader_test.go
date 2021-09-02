// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azblob

import (
	"context"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"math/rand"
	"net"
	"net/http"
	"time"
)

// Testings for RetryReader
// This reader return one byte through each Read call
type perByteReader struct {
	RandomBytes []byte // Random generated bytes

	byteCount              int // Bytes can be returned before EOF
	currentByteIndex       int // Bytes that have already been returned.
	doInjectError          bool
	doInjectErrorByteIndex int
	doInjectTimes          int
	injectedError          error

	// sleepDuraion and closeChannel are only use in "forced cancellation" tests
	sleepDuration time.Duration
	closeChannel  chan struct{}
}

func newPerByteReader(byteCount int) *perByteReader {
	perByteReader := perByteReader{
		byteCount:    byteCount,
		closeChannel: nil,
	}

	perByteReader.RandomBytes = make([]byte, byteCount)
	_, _ = rand.Read(perByteReader.RandomBytes)

	return &perByteReader
}

func newSingleUsePerByteReader(contents []byte) *perByteReader {
	perByteReader := perByteReader{
		byteCount:    len(contents),
		closeChannel: make(chan struct{}, 10),
	}

	perByteReader.RandomBytes = contents

	return &perByteReader
}

func (r *perByteReader) Read(b []byte) (n int, err error) {
	if r.doInjectError && r.doInjectErrorByteIndex == r.currentByteIndex && r.doInjectTimes > 0 {
		r.doInjectTimes--
		return 0, r.injectedError
	}

	if r.currentByteIndex < r.byteCount {
		n = copy(b, r.RandomBytes[r.currentByteIndex:r.currentByteIndex+1])
		r.currentByteIndex += n

		// simulate a delay, which may be successful or, if we're closed from another go-routine, may return an
		// error
		select {
		case <-r.closeChannel:
			return n, errors.New(ReadOnClosedBodyMessage)
		case <-time.After(r.sleepDuration):
			return n, nil
		}
	}

	return 0, io.EOF
}

func (r *perByteReader) Close() error {
	if r.closeChannel != nil {
		r.closeChannel <- struct{}{}
	}
	return nil
}

// Test normal retry succeed, note initial response not provided.
// Tests both with and without notification of failures
func (s *azblobUnrecordedTestSuite) TestRetryReaderReadWithRetry() {
	_assert := assert.New(s.T())
	// Test twice, the second time using the optional "logging"/notification callback for failed tries
	// We must test both with and without the callback, since be testing without
	// we are testing that it is, indeed, optional to provide the callback
	for _, logThisRun := range []bool{false, true} {

		// Extra setup for testing notification of failures (i.e. of unsuccessful tries)
		failureMethodNumCalls := 0
		failureWillRetryCount := 0
		failureLastReportedFailureCount := -1
		var failureLastReportedError error = nil
		failureMethod := func(failureCount int, lastError error, offset int64, count int64, willRetry bool) {
			failureMethodNumCalls++
			if willRetry {
				failureWillRetryCount++
			}
			failureLastReportedFailureCount = failureCount
			failureLastReportedError = lastError
		}

		// Main test setup
		byteCount := 1
		body := newPerByteReader(byteCount)
		body.doInjectError = true
		body.doInjectErrorByteIndex = 0
		body.doInjectTimes = 1
		body.injectedError = &net.DNSError{IsTemporary: true}

		getter := func(ctx context.Context, info HTTPGetterInfo) (*http.Response, error) {
			r := http.Response{}
			body.currentByteIndex = int(info.Offset)
			r.Body = body

			return &r, nil
		}

		httpGetterInfo := HTTPGetterInfo{Offset: 0, Count: int64(byteCount)}
		initResponse, err := getter(context.Background(), httpGetterInfo)
		_assert.Nil(err)

		rrOptions := RetryReaderOptions{MaxRetryRequests: 1}
		if logThisRun {
			rrOptions.NotifyFailedRead = failureMethod
		}
		retryReader := NewRetryReader(context.Background(), initResponse, httpGetterInfo, rrOptions, getter)

		// should fail and succeed through retry
		can := make([]byte, 1)
		n, err := retryReader.Read(can)
		_assert.Equal(n, 1)
		_assert.Nil(err)

		// check "logging", if it was enabled
		if logThisRun {
			// We only expect one failed try in this test
			// And the notification method is not called for successes
			_assert.Equal(failureMethodNumCalls, 1)           // this is the number of calls we counted
			_assert.Equal(failureWillRetryCount, 1)           // the sole failure was retried
			_assert.Equal(failureLastReportedFailureCount, 1) // this is the number of failures reported by the notification method
			_assert.NotNil(failureLastReportedError)
		}
		// should return EOF
		n, err = retryReader.Read(can)
		_assert.Equal(n, 0)
		_assert.Equal(err, io.EOF)
	}
}

// Test normal retry succeed, note initial response not provided.
// Tests both with and without notification of failures
func (s *azblobUnrecordedTestSuite) TestRetryReaderWithRetryIoUnexpectedEOF() {
	_assert := assert.New(s.T())
	// Test twice, the second time using the optional "logging"/notification callback for failed tries
	// We must test both with and without the callback, since be testing without
	// we are testing that it is, indeed, optional to provide the callback
	for _, logThisRun := range []bool{false, true} {

		// Extra setup for testing notification of failures (i.e. of unsuccessful tries)
		failureMethodNumCalls := 0
		failureWillRetryCount := 0
		failureLastReportedFailureCount := -1
		var failureLastReportedError error = nil
		failureMethod := func(failureCount int, lastError error, offset int64, count int64, willRetry bool) {
			failureMethodNumCalls++
			if willRetry {
				failureWillRetryCount++
			}
			failureLastReportedFailureCount = failureCount
			failureLastReportedError = lastError
		}

		// Main test setup
		byteCount := 1
		body := newPerByteReader(byteCount)
		body.doInjectError = true
		body.doInjectErrorByteIndex = 0
		body.doInjectTimes = 1
		body.injectedError = io.ErrUnexpectedEOF

		getter := func(ctx context.Context, info HTTPGetterInfo) (*http.Response, error) {
			r := http.Response{}
			body.currentByteIndex = int(info.Offset)
			r.Body = body

			return &r, nil
		}

		httpGetterInfo := HTTPGetterInfo{Offset: 0, Count: int64(byteCount)}
		initResponse, err := getter(context.Background(), httpGetterInfo)
		_assert.Nil(err)

		rrOptions := RetryReaderOptions{MaxRetryRequests: 1}
		if logThisRun {
			rrOptions.NotifyFailedRead = failureMethod
		}
		retryReader := NewRetryReader(context.Background(), initResponse, httpGetterInfo, rrOptions, getter)

		// should fail and succeed through retry
		can := make([]byte, 1)
		n, err := retryReader.Read(can)
		_assert.Equal(n, 1)
		_assert.Nil(err)

		// check "logging", if it was enabled
		if logThisRun {
			// We only expect one failed try in this test
			// And the notification method is not called for successes
			_assert.Equal(failureMethodNumCalls, 1)           // this is the number of calls we counted
			_assert.Equal(failureWillRetryCount, 1)           // the sole failure was retried
			_assert.Equal(failureLastReportedFailureCount, 1) // this is the number of failures reported by the notification method
			_assert.NotNil(failureLastReportedError)
		}
		// should return EOF
		n, err = retryReader.Read(can)
		_assert.Equal(n, 0)
		_assert.Equal(err, io.EOF)
	}
}

// Test normal retry fail as retry Count not enough.
func (s *azblobUnrecordedTestSuite) TestRetryReaderReadNegativeNormalFail() {
	_assert := assert.New(s.T())
	// Extra setup for testing notification of failures (i.e. of unsuccessful tries)
	failureMethodNumCalls := 0
	failureWillRetryCount := 0
	failureLastReportedFailureCount := -1
	var failureLastReportedError error = nil
	failureMethod := func(failureCount int, lastError error, offset int64, count int64, willRetry bool) {
		failureMethodNumCalls++
		if willRetry {
			failureWillRetryCount++
		}
		failureLastReportedFailureCount = failureCount
		failureLastReportedError = lastError
	}

	// Main test setup
	byteCount := 1
	body := newPerByteReader(byteCount)
	body.doInjectError = true
	body.doInjectErrorByteIndex = 0
	body.doInjectTimes = 2
	body.injectedError = &net.DNSError{IsTemporary: true}

	startResponse := http.Response{}
	startResponse.Body = body

	getter := func(ctx context.Context, info HTTPGetterInfo) (*http.Response, error) {
		r := http.Response{}
		body.currentByteIndex = int(info.Offset)
		r.Body = body

		return &r, nil
	}

	rrOptions := RetryReaderOptions{
		MaxRetryRequests: 1,
		NotifyFailedRead: failureMethod}
	retryReader := NewRetryReader(context.Background(), &startResponse, HTTPGetterInfo{Offset: 0, Count: int64(byteCount)}, rrOptions, getter)

	// should fail
	can := make([]byte, 1)
	n, err := retryReader.Read(can)
	_assert.Equal(n, 0)
	_assert.Equal(err, body.injectedError)

	// Check that we recieved the right notification callbacks
	// We only expect two failed tries in this test, but only one
	// of the would have had willRetry = true
	_assert.Equal(failureMethodNumCalls, 2)           // this is the number of calls we counted
	_assert.Equal(failureWillRetryCount, 1)           // only the first failure was retried
	_assert.Equal(failureLastReportedFailureCount, 2) // this is the number of failures reported by the notification method
	_assert.NotNil(failureLastReportedError)
}

// Test boundary case when Count equals to 0 and fail.
func (s *azblobUnrecordedTestSuite) TestRetryReaderReadCount0() {
	_assert := assert.New(s.T())
	byteCount := 1
	body := newPerByteReader(byteCount)
	body.doInjectError = true
	body.doInjectErrorByteIndex = 1
	body.doInjectTimes = 1
	body.injectedError = &net.DNSError{IsTemporary: true}

	startResponse := http.Response{}
	startResponse.Body = body

	getter := func(ctx context.Context, info HTTPGetterInfo) (*http.Response, error) {
		r := http.Response{}
		body.currentByteIndex = int(info.Offset)
		r.Body = body

		return &r, nil
	}

	retryReader := NewRetryReader(context.Background(), &startResponse, HTTPGetterInfo{Offset: 0, Count: int64(byteCount)}, RetryReaderOptions{MaxRetryRequests: 1}, getter)

	// should consume the only byte
	can := make([]byte, 1)
	n, err := retryReader.Read(can)
	_assert.Equal(n, 1)
	_assert.Nil(err)

	// should not read when Count=0, and should return EOF
	n, err = retryReader.Read(can)
	_assert.Equal(n, 0)
	_assert.Equal(err, io.EOF)
}

func (s *azblobUnrecordedTestSuite) TestRetryReaderReadNegativeNonRetriableError() {
	_assert := assert.New(s.T())
	byteCount := 1
	body := newPerByteReader(byteCount)
	body.doInjectError = true
	body.doInjectErrorByteIndex = 0
	body.doInjectTimes = 1
	body.injectedError = fmt.Errorf("not retriable error")

	startResponse := http.Response{}
	startResponse.Body = body

	getter := func(ctx context.Context, info HTTPGetterInfo) (*http.Response, error) {
		r := http.Response{}
		body.currentByteIndex = int(info.Offset)
		r.Body = body

		return &r, nil
	}

	retryReader := NewRetryReader(context.Background(), &startResponse, HTTPGetterInfo{Offset: 0, Count: int64(byteCount)}, RetryReaderOptions{MaxRetryRequests: 2}, getter)

	dest := make([]byte, 1)
	_, err := retryReader.Read(dest)
	_assert.Equal(err, body.injectedError)
}

// Test the case where we programmatically force a retry to happen, via closing the body early from another goroutine
// Unlike the retries orchestrated elsewhere in this test file, which simulate network failures for the
// purposes of unit testing, here we are testing the cancellation mechanism that is exposed to
// consumers of the API, to allow programmatic forcing of retries (e.g. if the consumer deems
// the read to be taking too long, they may force a retry in the hope of better performance next time).
func (s *azblobUnrecordedTestSuite) TestRetryReaderReadWithForcedRetry() {
	_assert := assert.New(s.T())
	for _, enableRetryOnEarlyClose := range []bool{false, true} {

		// use the notification callback, so we know that the retry really did happen
		failureMethodNumCalls := 0
		failureMethod := func(failureCount int, lastError error, offset int64, count int64, willRetry bool) {
			failureMethodNumCalls++
		}

		// Main test setup
		byteCount := 10 // so multiple passes through read loop will be required
		sleepDuration := 100 * time.Millisecond
		randBytes := make([]byte, byteCount)
		_, _ = rand.Read(randBytes)
		getter := func(ctx context.Context, info HTTPGetterInfo) (*http.Response, error) {
			body := newSingleUsePerByteReader(randBytes) // make new one every time, since we force closes in this test, and its unusable after a close
			body.sleepDuration = sleepDuration
			r := http.Response{}
			body.currentByteIndex = int(info.Offset)
			r.Body = body

			return &r, nil
		}

		httpGetterInfo := HTTPGetterInfo{Offset: 0, Count: int64(byteCount)}
		initResponse, err := getter(context.Background(), httpGetterInfo)
		_assert.Nil(err)

		rrOptions := RetryReaderOptions{MaxRetryRequests: 2, TreatEarlyCloseAsError: !enableRetryOnEarlyClose}
		rrOptions.NotifyFailedRead = failureMethod
		retryReader := NewRetryReader(context.Background(), initResponse, httpGetterInfo, rrOptions, getter)

		// set up timed cancellation from separate goroutine
		go func() {
			time.Sleep(sleepDuration * 5)
			err := retryReader.Close()
			if err != nil {
				return
			}
		}()

		// do the read (should fail, due to forced cancellation, and succeed through retry)
		output := make([]byte, byteCount)
		n, err := io.ReadFull(retryReader, output)
		if enableRetryOnEarlyClose {
			_assert.Equal(n, byteCount)
			_assert.Nil(err)
			_assert.EqualValues(output, randBytes)
			_assert.Equal(failureMethodNumCalls, 1) // assert that the cancellation did indeed happen
		} else {
			_assert.NotNil(err)
		}
	}
}
