package utils

import (
    "context"
    "encoding/csv"
    "os"
    "sort"
    "strconv"
    "time"
)

type Note struct {
    Timestamp time.Time `json:"timestamp"`
    UserID    int       `json:"user_id"`
    Note      string    `json:"note"`
}

func ReadRecentNotes(ctx context.Context, userID int, limit int) ([]Note, error) {
    file, err := os.Open("notes.csv")
    if err != nil {
        return nil, err
    }
    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        return nil, err
    }

    var notes []Note
    for _, record := range records {
        select {
        case <-ctx.Done():
            return nil, ctx.Err() // Handle context cancellation
        default:
        }

        userIDCsv, _ := strconv.Atoi(record[1])
        if userIDCsv == userID {
            timestamp, _ := strconv.ParseInt(record[0], 10, 64)
            note := Note{
                Timestamp: time.UnixMicro(timestamp),
                UserID:    userIDCsv,
                Note:      record[2],
            }
            notes = append(notes, note)
        }
    }

    // Sort notes by timestamp (this should be ascending already)
    sort.Slice(notes, func(i, j int) bool {
        return notes[i].Timestamp.Before(notes[j].Timestamp)
    })

    // Return only the most recent 'limit' notes
    if len(notes) > limit {
        return notes[len(notes)-limit:], nil
    }
    return notes, nil
}
