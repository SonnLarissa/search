package search

import (
	"context"
	"io/ioutil"
	"log"
	"strings"
	"sync"
)

//Describe searches result

type Result struct {
	Phrase  string
	Line    string
	LineNum int64
	ColNum  int64
}

//All - total searching
func All(ctx context.Context, phrase string, files []string) <-chan []Result {
	ch := make(chan []Result)
	wg := sync.WaitGroup{}
	//root := context.Background()
	ctx, cansel := context.WithCancel(ctx)

	for i := 0; i < len(files); i++ {
		wg.Add(1)

		go func(ctx context.Context, currentFile string, i int, ch chan<- []Result) {
			defer wg.Done()
			result, err := findAll(currentFile, phrase)

			if err != nil {
				log.Println(err)
			}

			if len(result) > 0 {
				ch <- result
			}

		}(ctx, files[i], i, ch)
	}

	go func() {
		defer close(ch)
		wg.Wait()
	}()

	cansel()
	return ch
}

func findAll(filePath string, phrase string) (res []Result, err error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println("ERROR, file not found")
		return res, err
	}
	arr := strings.Split(string(data), "\n")

	for i, str := range arr {
		ind := strings.Index(str, phrase)
		if ind > -1 {
			found := Result{
				Phrase:  phrase,
				Line:    str,
				LineNum: int64(i + 1),
				ColNum:  int64(ind) + 1,
			}
			res = append(res, found)
		}
	}
	return res, nil
}

//Any - search first phrase
func Any(ctx context.Context, phrase string, files []string) <-chan Result {
	ch := make(chan Result)
	wg := sync.WaitGroup{}
	results := Result{}
	//root := context.Background()
	ctx, cansel := context.WithCancel(ctx)

	for i := 0; i < len(files); i++ {

		data, err := ioutil.ReadFile(files[i])
		if err != nil {
			log.Println("file not found", err)
		}

		if strings.Contains(string(data), phrase) {
			parceRes, err := ParceForAny(string(data), phrase)
			if err != nil {
				log.Println(" parsing error")
			}
			if (Result{} != parceRes) {
				results = parceRes
				break
			}
		}

		wg.Add(1)

		go func(ctx context.Context, ch chan<- Result) {
			defer wg.Done()
			if (Result{}) != results {
				ch <- results
			}

		}(ctx, ch)
	}

	go func() {
		defer close(ch)
		wg.Wait()
	}()

	cansel()
	return ch
}

func ParceForAny(textString string, phrase string) (res Result, err error) {
	arr := strings.Split(string(textString), "\n")

	for i, str := range arr {
		ind := strings.Index(str, phrase)
		if ind > -1 {
			return  Result{
				Phrase:  phrase,
				Line:    str,
				LineNum: int64(i + 1),
				ColNum:  int64(ind) + 1,
			}, nil
		}
	}
	return res, nil
}
