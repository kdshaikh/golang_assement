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

func StreamVideo(sources []VideoSource, wg *sync.WaitGroup) {
	defer wg.Done()
	var segments []VideoSegment
	for _, source := range sources {
		for i := 1; i <= 3; i++ { // Fetch 3 segments from each source
			segment := source.FetchSegment(i)
			fmt.Printf("%s: %s fetched successfully\n", segment.source, segment.data)
			segments = append(segments, segment)
		}
	}
	fmt.Println("Video streaming completed using FanOut pattern")
	fmt.Println("Waiting for video segments to be aggregated...")
	aggregateVideoSegments(segments)
}

func aggregateVideoSegments(segments []VideoSegment) {
	var wg sync.WaitGroup
	aggregatedSegments := make(chan VideoSegment, len(segments))

	for _, segment := range segments {
		wg.Add(1)
		go func(segment VideoSegment) {
			defer wg.Done()
			aggregatedSegments <- segment
		}(segment)
	}

	go func() {
		wg.Wait()
		close(aggregatedSegments)
	}()

	for segment := range aggregatedSegments {
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

	sources := createVideoSources()
	fmt.Println("All sources->", sources)

	wg.Add(1)
	go StreamVideo(sources, &wg)

	wg.Wait()
}
