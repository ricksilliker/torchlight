package torchlight

import (
	"context"
	"github.com/sirupsen/logrus"
	"regexp"
	"strconv"
)

type FileReference struct {
	Namespace string `json:"namespace"`
	Filepath string `json:"source"`
	Depth int `json:"depth"`
	MayaScene string `json:"scene"`
	Loaded bool `json:"loaded"`
}

func GetAllReferences(ctx context.Context, sceneData, scenePath string) []FileReference {
	logger := ctx.Value("logger").(logrus.Logger)

	fileStr := sceneData
	fp := scenePath

	referencePattern, _ := regexp.Compile(`file.*-rfn.*\s*.*";`)
	matches := referencePattern.FindAllString(fileStr, -1)

	logger.WithField("num_matches", len(matches)).Info("Matches found")

	var sceneRefs []FileReference
	for _, x := range matches {
		sceneRefs = append(sceneRefs, FileReference{
			Namespace: GetReferenceNamespace(x),
			Filepath:  GetReferenceSourceFilePath(x),
			Depth:     GetReferenceDepth(x),
			MayaScene: fp,
			Loaded:    GetReferenceIsLoaded(x),
		})
	}

	return sceneRefs
}

func GetReferenceNamespace(refRequest string) string {
	referenceNSPattern, _ := regexp.Compile(`-ns "(?P<namespace>.*?)"`)
	ns := referenceNSPattern.FindStringSubmatch(refRequest)
	if len(ns) > 0 {
		return ns[1]
	}

	return ""
}

func GetReferenceSourceFilePath(refRequest string) string {
	referenceFilePathPattern, _ := regexp.Compile(`file.*-rfn.*\s*.*"(?P<filepath>.*)";`)
	fp := referenceFilePathPattern.FindStringSubmatch(refRequest)
	if len(fp) > 0 {
		return fp[1]
	}

	return ""
}

func GetReferenceDepth(refRequest string) int {
	referenceDepthPattern, _ := regexp.Compile(`-rdi (?P<depth>\d)`)

	depth := referenceDepthPattern.FindStringSubmatch(refRequest)
	var refDepthNum int
	refDepthNum = -1
	if len(depth) != 0 {
		num, err := strconv.Atoi(depth[1])
		if err != nil {
			refDepthNum = -1
		} else {
			refDepthNum = num
		}
	}

	return refDepthNum
}

func GetReferenceIsLoaded(refRequest string) bool {
	referenceDeferredPattern, _ := regexp.Compile(`-dr (?P<deferred>1)`)

	deferred := referenceDeferredPattern.FindStringSubmatch(refRequest)
	var loaded bool
	if len(deferred) == 0 {
		loaded = true
	}

	return loaded
}