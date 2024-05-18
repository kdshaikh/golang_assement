# Video Streaming Service

## Overview

This program demonstrates the use of concurrency patterns in Go to handle video streaming data across multiple sources. It utilizes the fanout pattern to concurrently fetch video segments from different sources and the fanin pattern to aggregate the segments for seamless playback.

## Features

- **Concurrent Video Fetching**:  Retrieves video segments concurrently from multiple sources.
- **Seamless Playback**: Aggregates video segments using the fanin pattern for uninterrupted streaming.

## Code Explanation

### Structures and Interfaces

#### VideoSegment

```go
type VideoSegment struct {
    segmentID   int
    source      string
    data        []byte
}
```

#### VideoSource

```go
type VideoSource interface {
    FetchSegment(segmentID int) VideoSegment
}
```

#### VideoStream

```go
type VideoStream struct {
    source string
}
```

### Function

#### *createVideoSources*
- Creates a list of video sources with different URLs.

#### *StreamVideo*
- Orchestrates the streaming of video segments from all sources and waits for all operations to complete.
- It uses FanOut Patern

#### *aggregateVideoSegments*
- Aggregates video segments from all sources for seamless playback.
- It uses FanIn Patern

### Exampl Output

```
all sources-> [{{youtube}} {{vimeo}} {{dailymotion}} {{twitch}}]
YouTube: Segment 1 fetched successfully
Vimeo: Segment 1 fetched successfully
Dailymotion: Segment 1 fetched successfully
Twitch: Segment 1 fetched successfully
Video streaming completed using FanOut pattern
Waiting for video segments to be aggregated...

Video playback started...
Video playback completed successfully.
```