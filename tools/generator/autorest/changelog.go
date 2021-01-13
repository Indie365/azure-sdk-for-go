package autorest

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Azure/azure-sdk-for-go/tools/apidiff/exports"
	"github.com/Azure/azure-sdk-for-go/tools/apidiff/report"
	"github.com/Azure/azure-sdk-for-go/tools/generator/autorest/model"
	"github.com/Azure/azure-sdk-for-go/tools/generator/changelog"
	"github.com/Azure/azure-sdk-for-go/tools/generator/utils"
)

type ChangelogContext interface {
	SDKRoot() string
	SDKCloneRoot() string
	SpecRoot() string
	SpecCommitHash() string
	CodeGenVersion() string
}

type changelogProcessor struct {
	ctx              ChangelogContext
	metadataLocation string
	readme           string
}

func NewChangelogProcessorFromContext(ctx ChangelogContext) *changelogProcessor {
	return &changelogProcessor{
		ctx: ctx,
	}
}

func (p *changelogProcessor) WithLocation(metadataLocation string) *changelogProcessor {
	p.metadataLocation = metadataLocation
	return p
}

func (p *changelogProcessor) WithReadme(readme string) *changelogProcessor {
	// make sure the readme here is a relative path to the root of spec
	readme = utils.NormalizePath(readme)
	root := utils.NormalizePath(p.ctx.SpecRoot())
	if filepath.IsAbs(readme) {
		readme = strings.TrimPrefix(readme, root)
	}
	p.readme = readme
	return p
}

type ChangelogResult struct {
	PackageName        string
	PackagePath        string
	GenerationMetadata changelog.GenerationMetadata
	Changelog          changelog.Changelog
}

type ChangelogProcessError struct {
	Errors []error
}

func (e *ChangelogProcessError) Error() string {
	return fmt.Sprintf("total %d error(s) during processing changelog: %+v", len(e.Errors), e.Errors)
}

type changelogErrorBuilder struct {
	errors []error
}

func (b *changelogErrorBuilder) add(err error) {
	b.errors = append(b.errors, err)
}

func (b *changelogErrorBuilder) build() error {
	if len(b.errors) == 0 {
		return nil
	}
	return &ChangelogProcessError{
		Errors: b.errors,
	}
}

// Process generates the changelogs using the input metadata map.
// Please ensure the input metadata map does not contain any package that is not under the sdk root, otherwise this might give weird results.
func (p *changelogProcessor) Process(metadataMap map[string]model.Metadata) ([]ChangelogResult, error) {
	builder := changelogErrorBuilder{}
	var results []ChangelogResult
	for tag, metadata := range metadataMap {
		outputFolder := filepath.Clean(metadata.PackagePath())
		// format the package before getting the changelog
		if err := FormatPackage(outputFolder); err != nil {
			builder.add(err)
			continue
		}
		// use the relative path to the sdk root as package name
		packageName, err := filepath.Rel(p.ctx.SDKRoot(), outputFolder)
		if err != nil {
			builder.add(err)
			continue
		}
		// normalize the package name
		packageName = utils.NormalizePath(packageName)
		result, err := p.GenerateChangelog(packageName, tag)
		if err != nil {
			builder.add(err)
			continue
		}
		result.PackagePath = outputFolder
		results = append(results, *result)
	}
	return results, builder.build()
}

func (p *changelogProcessor) GenerateChangelog(packageName, tag string) (*ChangelogResult, error) {
	lhs, err := getExportsForPackage(filepath.Join(p.ctx.SDKCloneRoot(), packageName))
	if err != nil {
		return nil, fmt.Errorf("failed to get exports from package '%s' in the clone '%s': %+v", packageName, p.ctx.SDKCloneRoot(), err)
	}
	rhs, err := getExportsForPackage(filepath.Join(p.ctx.SDKRoot(), packageName))
	if err != nil {
		return nil, fmt.Errorf("failed to get exports from package '%s' in the sdk '%s': %+v", packageName, p.ctx.SDKRoot(), err)
	}
	r, err := getChangelogForPackage(packageName, lhs, rhs)
	if err != nil {
		return nil, err
	}
	return &ChangelogResult{
		PackageName: packageName,
		GenerationMetadata: changelog.GenerationMetadata{
			CommitHash:     p.ctx.SpecCommitHash(),
			Readme:         p.readme,
			Tag:            tag,
			CodeGenVersion: p.ctx.CodeGenVersion(),
		},
		Changelog: *r,
	}, nil
}

func getExportsForPackage(dir string) (*exports.Content, error) {
	// The function exports.Get does not handle the circumstance that the package does not exist
	// therefore we have to check if it exists and if not exit early to ensure we do not return an error
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return nil, nil
	}
	exp, err := exports.Get(dir)
	if err != nil {
		return nil, err
	}
	return &exp, nil
}

func getChangelogForPackage(pkgDir string, lhs, rhs *exports.Content) (*changelog.Changelog, error) {
	if lhs == nil && rhs == nil {
		return nil, fmt.Errorf("this package does not exist even after the generation, this should never happen")
	}
	if lhs == nil {
		// the package does not exist before the generation: this is a new package
		return &changelog.Changelog{
			PackageName: pkgDir,
			NewPackage:  true,
		}, nil
	}
	if rhs == nil {
		// the package no longer exists after the generation: this package was removed
		return &changelog.Changelog{
			PackageName:    pkgDir,
			RemovedPackage: true,
		}, nil
	}
	// lhs and rhs are both non-nil
	p := report.Generate(*lhs, *rhs, nil)
	return &changelog.Changelog{
		PackageName: pkgDir,
		Modified:    &p,
	}, nil
}
