package main

import (
	"fmt"
	"sync"
	"time"
)

type VideoSegment struct {
	segmentID int
	source    string
	data      []byte
}

type VideoSource interface {
	FetchSegment(segmentID int) VideoSegment
}

type VideoStream struct {
	source string
}

func (v VideoStream) FetchSegment(segmentID int) VideoSegment {
	// Simulate fetching a video segment
	time.Sleep(time.Duration(segmentID) * time.Second)
	return VideoSegment{
		segmentID: segmentID,
		source:    v.source,
		data:      []byte(fmt.Sprintf("Segment %d data from %s", segmentID, v.source)),
	}
}

func createVideoSources() []VideoSource {
	videoSources := []string{"YouTube", "Vimeo", "Dailymotion", "Twitch"}
	var sources []VideoSource
	for _, source := range videoSources {
		sources = append(sources, VideoStream{source: source})
	}
	return sources
}

func StreamVideo(sources []VideoSource, wg *sync.WaitGroup, segmentChan chan<- VideoSegment) {
	defer wg.Done()
	for _, source := range sources {
		for i := 1; i <= 3; i++ { // Fetch 3 segments from each source
			segment := source.FetchSegment(i)
			fmt.Printf("%s: %s fetched successfully\n", segment.source, segment.data)
			segmentChan <- segment
		}
	}
}

func aggregateVideoSegments(segmentChan <-chan VideoSegment, wg *sync.WaitGroup) {
	defer wg.Done()
	for segment := range segmentChan {
		fmt.Printf("Received segment %d from %s\n", segment.segmentID, segment.source)
		// Process the received segment
	}

	fmt.Println("Video playback started...")
	// Simulate video playback
	time.Sleep(5 * time.Second)
	fmt.Println("Video playback completed successfully.")
}

func main() {
	var wg sync.WaitGroup
	var cg sync.WaitGroup

	sources := createVideoSources()
	fmt.Println("All sources->", sources)

	// Channel for collecting video segments
	segmentChan := make(chan VideoSegment)

	// Fan-out: Start goroutines to fetch segments from each source concurrently
	wg.Add(len(sources))
	for _, source := range sources {
		go StreamVideo([]VideoSource{source}, &wg, segmentChan)
	}

	// Fan-in: Start goroutine to aggregate segments
	cg.Add(1)
	go aggregateVideoSegments(segmentChan, &cg)

	// Wait for all goroutines to finish
	wg.Wait()
	close(segmentChan)
	cg.Wait()
}
