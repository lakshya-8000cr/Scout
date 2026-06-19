package cmd

import (
    "fmt"
    "strings"

    "github.com/fatih/color"
    "github.com/spf13/cobra"
)

var podsCmd = &cobra.Command{
    Use:   "pods",
    Short: "Show pods in a pretty format",
    Run: func(cmd *cobra.Command, args []string) {
        pods := []podRow{
            {"READY", "api-server", "Running", 0},
            {"READY", "postgres", "Running", 0},
            {"WARN", "redis", "Pending", 0},
            {"ERROR", "worker", "CrashLoopBackOff", 4},
        }

        printPodsTable(pods)
    },
}

func init() {
    rootCmd.AddCommand(podsCmd)
}


type podRow struct {
    label    string
    name     string
    status   string
    restarts int
}


func printPodsTable(pods []podRow) {
    width := 62

    fmt.Println()
    printTopBorder(width)
    printHeaderRow(width)
    printDivider(width)

    for _, p := range pods {
        printPodRow(p, width)
    }

    printBottomBorder(width)

    printSummary(pods)
    fmt.Println()
}

func printTopBorder(w int) {
    dim := color.New(color.Faint)
    dim.Printf("  ╭%s╮\n", strings.Repeat("─", w))
}

func printBottomBorder(w int) {
    dim := color.New(color.Faint)
    dim.Printf("  ╰%s╯\n", strings.Repeat("─", w))
}

func printDivider(w int) {
    dim := color.New(color.Faint)
    dim.Printf("  ├%s┤\n", strings.Repeat("─", w))
}

func printHeaderRow(w int) {
    color.New(color.Faint).Print("  │  ")
    color.New(color.BgHiBlue, color.FgWhite, color.Bold).Print(" SCOUT ")
    color.New(color.Faint).Print("  ")
    color.New(color.Bold, color.FgWhite).Print("Kubernetes Pod Status")
    
   
    color.New(color.Faint).Printf("%-*s│\n", w-32, "")
}

func printPodRow(p podRow, w int) {
    _ = w
    color.New(color.Faint).Print("  │  ")

    printBadge(p.label)

    color.New(color.FgWhite).Printf("  %-22s", p.name)


    printStatusText(p.status)

        if p.restarts > 0 {
        color.New(color.FgYellow).Printf("  restarts=%-3d", p.restarts)
    } else {
        color.New(color.Faint).Printf("  restarts=%-3d", p.restarts)
    }

    color.New(color.Faint).Print(" │\n")
}

func printBadge(label string) {
    switch label {
    case "READY":
        color.New(color.BgGreen, color.FgBlack, color.Bold).Printf(" %-5s ", label)
    case "WARN":
        color.New(color.BgYellow, color.FgBlack, color.Bold).Printf(" %-5s ", label)
    case "ERROR":
        color.New(color.BgRed, color.FgWhite, color.Bold).Printf(" %-5s ", label)
    default:
        color.New(color.BgWhite, color.FgBlack, color.Bold).Printf(" %-5s ", label)
    }
}

func printStatusText(status string) {
    switch status {
    case "Running":
        color.New(color.FgGreen).Printf("%-18s", status)
    case "Pending":
        color.New(color.FgYellow).Printf("%-18s", status)
    case "CrashLoopBackOff", "Error", "OOMKilled":
        color.New(color.FgRed).Printf("%-18s", status)
    default:
        color.New(color.Faint).Printf("%-18s", status)
    }
}

func printSummary(pods []podRow) {
    ready, warn, errs := 0, 0, 0
    for _, p := range pods {
        switch p.label {
        case "READY":
            ready++
        case "WARN":
            warn++
        case "ERROR":
            errs++
        }
    }

    color.New(color.Faint).Print("  ")
    color.New(color.FgGreen).Printf("● %d running", ready)
    color.New(color.Faint).Print("  ")
    color.New(color.FgYellow).Printf("● %d pending", warn)
    color.New(color.Faint).Print("  ")
    color.New(color.FgRed).Printf("● %d error", errs)
    color.New(color.Faint).Printf("  — %d pods total\n", len(pods))
}