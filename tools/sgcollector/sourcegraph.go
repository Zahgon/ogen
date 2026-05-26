package main

import (
	"context"
	"io"

	"github.com/go-faster/jx"
	"go.uber.org/zap/zapcore"
)

type (
	Repository struct {
		Name string `json:"name"`
	}

	ExternalURL struct {
		URL         string `json:"url"`
		ServiceKind string `json:"serviceKind"`
	}

	File struct {
		Name         string        `json:"name"`
		Size         int           `json:"size"`
		Path         string        `json:"path"`
		ByteSize     uint64        `json:"byteSize"`
		Content      string        `json:"content"`
		CanonicalURL string        `json:"canonicalURL"`
		ExternalURLs []ExternalURL `json:"externalURLs"`
	}

	FileMatch struct {
		Typename   string     `json:"__typename"`
		Repository Repository `json:"repository"`
		File       File       `json:"file"`
	}

	Alert struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	QueryVariables struct {
		Query string `json:"query"`
	}

	GraphQLQuery struct {
		Query     string `json:"query"`
		Variables QueryVariables
	}
)

func (m FileMatch) MarshalLogObject(e zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

func (m FileMatch) Link() string { _ = "STUB: not implemented"; return "" }

const graphQLQuery = `query ($query: String!) {
  search(query: $query, version: V2, patternType: regexp) {
    results {
      results {
        __typename
        ... on FileMatch {
          ...FileMatchFields
        }
      }
      limitHit
      matchCount
      elapsedMilliseconds
      ...SearchResultsAlertFields
    }
  }
}

fragment FileMatchFields on FileMatch {
  repository {
    name
  }
  file {
    name
    path
    canonicalURL
    externalURLs {
      serviceKind
      url
    }
    byteSize
    content
  }
}


fragment SearchResultsAlertFields on SearchResults {
  alert {
    title
    description
    proposedQueries {
      description
      query
    }
  }
}
`

func search(ctx context.Context, query string, cb func(FileMatch) error) (rerr error) {
	_ = "STUB: not implemented"
	return nil
}

func querySourcegraph(ctx context.Context, q GraphQLQuery, cb func(d *jx.Decoder, key []byte) error) error {
	_ = "STUB: not implemented"
	// Handling of the response may take a while, so we save the response to avoid timeouts.
	return nil
}

func sendSourcegraph(ctx context.Context, q GraphQLQuery, out io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

//#nosec G704
