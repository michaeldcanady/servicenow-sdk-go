package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type TestEvent struct {
	Time    string
	Action  string
	Package string
	Test    string
	Output  string
	Elapsed float64
}

type PackageResult struct {
	Name        string
	Passed      int
	Failed      int
	Skipped     int
	Elapsed     float64
	FailedTests []string
}

func main() {
	results := parseTestEvents(os.Stdin)

	printReportTable(results)
	printFailedTests(results)
}

func parseTestEvents(input *os.File) map[string]*PackageResult {
	results := make(map[string]*PackageResult)
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		var event TestEvent
		if err := json.Unmarshal(scanner.Bytes(), &event); err != nil {
			continue
		}

		if _, ok := results[event.Package]; !ok {
			results[event.Package] = &PackageResult{Name: event.Package}
		}

		res := results[event.Package]
		updatePackageResult(res, &event)
	}
	return results
}

func updatePackageResult(res *PackageResult, event *TestEvent) {
	switch event.Action {
	case "pass":
		if event.Test != "" {
			res.Passed++
		} else {
			res.Elapsed = event.Elapsed
		}
	case "fail":
		if event.Test != "" {
			res.Failed++
			res.FailedTests = append(res.FailedTests, event.Test)
		}
	case "skip":
		res.Skipped++
	}
}

func printReportTable(results map[string]*PackageResult) {
	fmt.Println("### 🧪 Go Test Report")
	fmt.Println("| Status  | Elapsed | Package | Cover | Pass | Fail | Skip |")
	fmt.Println("|---------|---------|---------|-------|------|------|------|")

	packages := getSortedPackageNames(results)

	for _, pkgName := range packages {
		res := results[pkgName]
		status := "🟢 PASS"
		if res.Failed > 0 {
			status = "🔴 FAIL"
		}
		fmt.Printf("| %s | %.2fs | %s | -- | %d | %d | %d |\n", status, res.Elapsed, res.Name, res.Passed, res.Failed, res.Skipped)
	}
}

func printFailedTests(results map[string]*PackageResult) {
	packages := getSortedPackageNames(results)
	hasFailures := false
	for _, res := range results {
		if res.Failed > 0 {
			hasFailures = true
			break
		}
	}

	if hasFailures {
		fmt.Println("\n### ❌ Failed Tests")
		fmt.Println("<details>")
		fmt.Println("<summary>Click to expand failed tests</summary>")
		for _, pkgName := range packages {
			res := results[pkgName]
			if res.Failed > 0 {
				fmt.Printf("\n## 🔴 FAIL • %s\n\n```\n", res.Name)
				for _, t := range res.FailedTests {
					fmt.Printf("--- FAIL: %s\n", t)
				}
				fmt.Println("```")
			}
		}
		fmt.Println("</details>")
	}
}

func getSortedPackageNames(results map[string]*PackageResult) []string {
	var packages []string
	for pkg := range results {
		packages = append(packages, pkg)
	}
	sort.Strings(packages)
	return packages
}
