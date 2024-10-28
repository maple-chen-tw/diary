package controller

import (
	models "diary/model"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// 假設的全局變量，用於存儲日記條目
var diaries []models.UserDiary

// GetAllDiary 獲取所有日記條目
func GetAllDiary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(diaries)
}

// CreateDiary 創建新的日記條目
func CreateDiary(w http.ResponseWriter, r *http.Request) {
	var newDiary models.UserDiary
	if err := json.NewDecoder(r.Body).Decode(&newDiary); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	diaries = append(diaries, newDiary)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newDiary)
}

// SaveDiary 更新日記條目
func SaveDiary(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]          // 獲取 id 字串
	id, err := strconv.Atoi(idStr) // 將字串轉換為整數
	if err != nil {
		http.Error(w, "Invalid diary ID", http.StatusBadRequest)
		return
	}

	for i, diary := range diaries {
		if diary.ID == id { // 使用整數進行比較
			if err := json.NewDecoder(r.Body).Decode(&diaries[i]); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			diaries[i].UpdateAt = time.Now()      // 更新最後修改時間
			json.NewEncoder(w).Encode(diaries[i]) // 返回更新後的日記條目
			return
		}
	}
	http.Error(w, "Diary not found", http.StatusNotFound) // 找不到日記條目時返回錯誤
}

// UndoDiary 恢復日記條目（此處可以根據具體需求實現）
func UndoDiary(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	// 實現恢復邏輯，例如標記為未完成
	for _, diary := range diaries {
		if diary.ID == params["id"] {
			// 進行相應的恢復處理
			json.NewEncoder(w).Encode(diary) // 返回相應的日記條目
			return
		}
	}
	http.Error(w, "Diary not found", http.StatusNotFound)
}

// DeleteDiary 刪除指定的日記條目
func DeleteDiary(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, diary := range diaries {
		if diary.ID == params["id"] {
			diaries = append(diaries[:i], diaries[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Diary not found", http.StatusNotFound)
}

// DeleteAllDiary 刪除所有日記條目
func DeleteAllDiary(w http.ResponseWriter, r *http.Request) {
	diaries = []models.UserDiary{} // 清空日記列表
	w.WriteHeader(http.StatusNoContent)
}
