package utils

import (
    "fmt"
    "math/rand"
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

func CreateNotesFile(numNotes int, numUsers int) {
    // Open or create notes.csv file
    file, err := os.Create("notes.csv")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close()

    // Set up random seed
    rand.Seed(time.Now().UnixNano())

    // Create CSV header
    file.WriteString("timestamp,user_id,note\n")

    // Generate rows with random data
    for i := 0; i < numNotes; i++ {
        // Generate random timestamp (in microseconds)
        timestamp := time.Now().UnixMicro() + int64(rand.Intn(1000000))

        // Generate random user_id between 1 and numUsers
        userID := rand.Intn(numUsers) + 1

        // Generate a random note
        note := "This is note number " + strconv.Itoa(i+1)

        // Write row to file
        file.WriteString(fmt.Sprintf("%d,%d,%s\n", timestamp, userID, note))
    }

    fmt.Printf("Generated %d rows in notes.csv\n", numNotes)
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
