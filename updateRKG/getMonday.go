package main

import (
    "time"
)

func getThisWeekMonday() time.Time {
    now := time.Now()
    
    // 現在の曜日を取得（0 = 日曜日, 1 = 月曜日, ..., 6 = 土曜日）
    weekday := now.Weekday()
    
    // 月曜日までの日数を計算
    var daysUntilMonday int
    if weekday == time.Sunday {
        daysUntilMonday = -6 // 日曜日の場合、前の週の月曜日まで戻る
    } else {
        daysUntilMonday = int(time.Monday - weekday)
    }
    
    // 今週の月曜日の日付を取得
    monday := now.AddDate(0, 0, daysUntilMonday)
    
    // 時刻部分を0時0分0秒にリセット
    return time.Date(monday.Year(), monday.Month(), monday.Day(), 0, 0, 0, 0, monday.Location())
}
