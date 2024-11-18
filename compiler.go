package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

func downloadFile(url string, filepath string) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func compileAndRunCSharp() error { // create a temp.cs failsafe script and compile it using go
	csharpCode := `
using System;
using System.Net;

class Program
{
    static void Main()
    {
        string url = "http://example.com/package.zip"; // URL of the package
        string destination = "package.zip"; // Local file name

        using (WebClient client = new WebClient())
        {
            try
            {
                client.DownloadFile(url, destination);
                Console.WriteLine("Package downloaded successfully.");
            }
            catch (Exception e)
            {
                Console.WriteLine("Error downloading package: " + e.Message);
            }
        }
    }
}
`
	// Write the C# code to a temporary file
	csharpFile := "temp.cs"
	err := os.WriteFile(csharpFile, []byte(csharpCode), 0644)
	if err != nil {
		return err
	}
	defer os.Remove(csharpFile) // Clean up

	// Compile the C# script
	cmd := exec.Command("csc", csharpFile)
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("compilation failed: %w", err)
	}

	// Run the compiled C# executable
	csharpExe := "temp.exe"
	cmd = exec.Command(csharpExe)
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("execution failed: %w", err)
	}

	return nil
}

func main() {
	url := "http://example.com/package.zip" // replace with the url of the web server you hosted on codespace
	err := downloadFile(url, "package.zip")
	if err != nil {
		fmt.Println("Error downloading package:", err)
		fmt.Println("Attempting to compile...")
		if err := compileAndRunCSharp(); err != nil {
			fmt.Println("Failed to run C# script:", err)
		}
	} else {
		fmt.Println("Package downloaded successfully.")
	}
}
