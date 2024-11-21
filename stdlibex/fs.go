package stdlibex

import (
	"bufio"
	"errors"
	"io/fs"
	"os"
	"strings"
	"syscall"
)

func DirExists(dir string) (exists bool, err error) {
	// TODO Differentiate that this is not a file
	return EntryExists(dir)
}
func FileExists(file string) (exists bool, err error) {
	// TODO Differentiate that this is not a dir
	return EntryExists(file)
}
func EntryExists(entry string) (exists bool, err error) {
	_, err = os.Stat(entry)
	if errors.Is(err, fs.ErrNotExist) {
		err = nil
		exists = false
		goto end
	}
	if err != nil {
		goto end
	}
	exists = true
end:
	return exists, err
}

func FileReadLines(filename string) (lines []string, err error) {
	var scanner *bufio.Scanner
	var file *os.File

	lines = make([]string, 0)
	file, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		goto end
	}
	defer MustClose(file)
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			continue
		}
		lines = append(lines, text)
	}
	err = scanner.Err()
end:
	return lines, err
}

func FileAppendLine(line string, filename string) (err error) {
	var f *os.File
	f, err = os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		goto end
	}
	defer func() {
		closeErr := f.Close()
		if closeErr == nil {
			return
		}
		if err == nil {
			err = closeErr
			return
		}
		err = errors.Join(err, closeErr)
	}()
	_, err = f.WriteString(strings.TrimSpace(line) + "\n")
end:
	return err
}

func FileSaveLines(filename string, lines []string) (err error) {
	err = os.Remove(filename)
	if err != nil && !os.IsNotExist(err) {
		goto end
	}
	err = FileAppendLine(strings.Join(lines, "\n"), filename)
end:
	return err
}

func FileRemoveLine(line, filename string) (err error) {
	var lines []string
	lines, err = FileReadLines(filename)
	n := 0
	for i := range lines {
		lines[n] = lines[i]
		if strings.TrimSpace(line) == lines[i] {
			n++
		}
	}
	if len(lines) == 0 {
		err = os.Remove(filename)
	} else {
		err = FileSaveLines(filename, lines[:n])
	}
	return err
}

func EnsureDir(dir string) (err error) {
	err = os.Mkdir(dir, os.ModePerm)
	if errors.Is(err, syscall.EEXIST) {
		err = nil
		goto end
	}
end:
	return err
}

func EnsureFileRemoved(file string) (err error) {
	err = os.Remove(file)
	if errors.Is(err, syscall.ENOENT) {
		err = nil
		goto end
	}
end:
	return err
}

func TouchFile(path string) (err error) {
	var fp *os.File
	fp, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if fp != nil {
		err = fp.Close()
	}
	return err
}
