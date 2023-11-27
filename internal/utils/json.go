package utils

import (
    "runtime"
    "encoding/json"
)

type Maintainer struct {
    Email string `json:"email"`
    Website string `json:"website"`
}

type Project struct {
    Name string `json:"name"`
    Description string `json:"description"`
    Email string `json:"email"`
    Website string `json:"website"`
    Repository string `json:"repository"`
    Version string `json:"version"`
}

type Response struct {
    Versions []string `json:"versions"`
    Language string `json:"language"`
    MemoryUsage uint64 `json:"memoryUsage"`
    Maintainer Maintainer `json:"maintainer"`
    Project Project `json:"project"`
}

func calcMem() uint64 {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    return m.Alloc / 1024 / 1024
}

func GetJson() []byte {
    jsonResp := Response{
        Versions: []string{"v3"},
        Language: "Go",
        MemoryUsage: calcMem(),
        Maintainer: Maintainer{
            Email: "support@rubynetwork.tech",
            Website: "https://rubynetwork.tech",
        },
        Project: Project{
            Name: "bare-go",
            Description: "A Bare Server in GoLang",
            Email: "support@rubynetwork.tech",
            Website: "https://rubynetwork.tech",
            Repository: "httos://github.com/ruby-network/bare-go",
            Version: "v0.0.1",
        },
    }
    prettyJson, _ := json.MarshalIndent(jsonResp, "", "  ")
    return prettyJson
}
