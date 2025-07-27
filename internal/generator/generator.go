package generator

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

// CreateRandomFiles creates random text files at specified depth with random names
func CreateRandomFiles(baseDir string, depth, count int) error {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Convert to absolute path
	absBaseDir, err := filepath.Abs(baseDir)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %v", err)
	}

	// Create directory structure if needed
	if depth > 1 {
		err := createDirectoryStructure(absBaseDir, depth)
		if err != nil {
			return fmt.Errorf("failed to create directory structure: %v", err)
		}
	}

	// Create files at the specified depth
	targetDir := getTargetDirectory(absBaseDir, depth)

	for i := 0; i < count; i++ {
		filename := generateRandomFilename() + ".txt"
		filePath := filepath.Join(targetDir, filename)

		content := generateRandomContent()

		err := os.WriteFile(filePath, []byte(content), 0644)
		if err != nil {
			return fmt.Errorf("failed to create file %s: %v", filePath, err)
		}

		fmt.Printf("Created: %s\n", filePath)
	}

	return nil
}

// createDirectoryStructure creates nested directories up to the specified depth
func createDirectoryStructure(baseDir string, depth int) error {
	path := baseDir
	for i := 1; i < depth; i++ {
		dirName := fmt.Sprintf("level_%d", i)
		path = filepath.Join(path, dirName)

		err := os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

// getTargetDirectory returns the target directory at the specified depth
func getTargetDirectory(baseDir string, depth int) string {
	if depth == 1 {
		return baseDir
	}

	path := baseDir
	for i := 1; i < depth; i++ {
		dirName := fmt.Sprintf("level_%d", i)
		path = filepath.Join(path, dirName)
	}
	return path
}

// generateRandomFilename generates a random filename
func generateRandomFilename() string {
	adjectives := []string{"quick", "lazy", "happy", "sad", "big", "small", "fast", "slow", "bright", "dark"}
	nouns := []string{"cat", "dog", "bird", "fish", "tree", "rock", "star", "moon", "sun", "cloud"}

	adj := adjectives[rand.Intn(len(adjectives))]
	noun := nouns[rand.Intn(len(nouns))]
	num := rand.Intn(1000)

	return fmt.Sprintf("%s_%s_%d", adj, noun, num)
}

// generateRandomContent generates random content for the file
func generateRandomContent() string {
	sentences := []string{
		"This is a randomly generated file.",
		"The quick brown fox jumps over the lazy dog.",
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
		"Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		"Ut enim ad minim veniam, quis nostrud exercitation ullamco.",
		"Duis aute irure dolor in reprehenderit in voluptate velit esse.",
		"Excepteur sint occaecat cupidatat non proident, sunt in culpa.",
	}

	content := fmt.Sprintf("Generated at: %s\n\n", time.Now().Format(time.RFC3339))

	// Add 3-7 random sentences
	numSentences := 3 + rand.Intn(5)
	for i := 0; i < numSentences; i++ {
		sentence := sentences[rand.Intn(len(sentences))]
		content += sentence + "\n"
	}

	return content
}
