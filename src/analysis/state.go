package analysis

import (
	"fmt"
	"math"
	"strings"

	"github.com/brunobmello25/educationalsp/src/lsp"
)

type State struct {
	Documents map[string]string
}

func NewState() State {
	return State{
		Documents: make(map[string]string),
	}
}

func getDiagnosticsForFile(text string) []lsp.Diagnostic {
	diagnostics := []lsp.Diagnostic{}

	for row, line := range strings.Split(text, "\n") {
		if strings.Contains(line, "VS Code") {
			idx := strings.Index(line, "VS Code")
			diagnostics = append(diagnostics, lsp.Diagnostic{
				Range:    LineRange(row, idx, idx+len("VS Code")),
				Severity: 1,
				Source:   "Common Sense",
				Message:  "Please make sure we don't curse (use VS Code)",
			})
		}

		if strings.Contains(line, "Neovim") {
			idx := strings.Index(line, "Neovim")
			diagnostics = append(diagnostics, lsp.Diagnostic{
				Range:    LineRange(row, idx, idx+len("Neovim")),
				Severity: 4,
				Source:   "Common Sense",
				Message:  "Great choice!",
			})
		}
	}

	return diagnostics
}

func (s *State) OpenDocument(uri string, text string) []lsp.Diagnostic {
	s.Documents[uri] = text

	return getDiagnosticsForFile(text)
}

func (s *State) UpdateDocument(uri string, text string) []lsp.Diagnostic {
	s.Documents[uri] = text

	return getDiagnosticsForFile(text)
}

func (s *State) Hover(id int, uri string, position lsp.Position) lsp.HoverResponse {
	document := s.Documents[uri]

	response := lsp.HoverResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.HoverResult{
			Contents: fmt.Sprintf("File: %s, characters: %d", uri, len(document)),
		},
	}

	return response
}

func (s *State) Definition(id int, uri string, position lsp.Position) lsp.DefinitionResponse {
	response := lsp.DefinitionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.Location{
			URI: uri,
			Range: lsp.Range{
				Start: lsp.Position{
					Line: int(math.Max(0, float64(position.Line-1))), Character: 0,
				},
				End: lsp.Position{
					Line:      int(math.Max(0, float64(position.Line-1))),
					Character: 0,
				},
			},
		},
	}

	return response
}

func (s *State) TextDocumentCodeAction(id int, uri string, position lsp.Position) lsp.TextDocumentCodeActionResponse {
	text := s.Documents[uri]

	actions := []lsp.CodeAction{}
	for row, line := range strings.Split(text, "\n") {
		idx := strings.Index(line, "VS Code")
		if idx >= 0 {
			replaceChange := map[string][]lsp.TextEdit{}
			replaceChange[uri] = []lsp.TextEdit{
				{
					Range:   LineRange(row, idx, idx+len("VS Code")),
					NewText: "Neovim",
				},
			}

			actions = append(actions, lsp.CodeAction{
				Title: "Replace VS C*de with Neovim",
				Edit: &lsp.WorkspaceEdit{
					Changes: replaceChange,
				},
			})

			censorChange := map[string][]lsp.TextEdit{}
			censorChange[uri] = []lsp.TextEdit{
				{
					Range:   LineRange(row, idx, idx+len("VS Code")),
					NewText: "VS C*de",
				},
			}

			actions = append(actions, lsp.CodeAction{
				Title: "Censor VS C*de",
				Edit: &lsp.WorkspaceEdit{
					Changes: censorChange,
				},
			})
		}
	}

	response := lsp.TextDocumentCodeActionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: actions,
	}

	return response
}

func LineRange(row int, start int, end int) lsp.Range {
	return lsp.Range{
		Start: lsp.Position{
			Line:      row,
			Character: start,
		},
		End: lsp.Position{
			Line:      row,
			Character: end,
		},
	}
}

func (s *State) TextDocumentCompletion(id int) lsp.CompletionResponse {
	items := []lsp.CompletionItem{
		{
			Label:         "Neovim (BTW)",
			Detail:        "Very cool editor",
			Documentation: "I like it",
		},
	}

	response := lsp.CompletionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: items,
	}

	return response
}
