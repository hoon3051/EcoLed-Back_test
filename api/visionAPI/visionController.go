package visionapi

import (
    "sort"
    "context"
    "net/http"

    vision "cloud.google.com/go/vision/apiv1"
    "github.com/gin-gonic/gin"
)

var filterKeywords = map[string]bool{
    "Plastic": true,
    "Paper":   true,
    "Glass":   true,
    "Aluminium": true,
    "Metal": true,
    "Steel":true,
    "Wood":true,
    "Battery": true,
    // 필요한 추가 키워드를 여기에 추가
}


func HandleFileUpload(c *gin.Context) {
    file, err := c.FormFile("file")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to form file",
		})
        return
    }
    
    ctx := context.Background()

    // 클라이언트 생성
    client, err := vision.NewImageAnnotatorClient(ctx)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create client",
		})
        return
    }
    defer client.Close()

    // 이미지 파일 열기
    f, err := file.Open()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to open image",
		})
        return
    }
    defer f.Close()

    image, err := vision.NewImageFromReader(f)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to read image",
		})
        return
    }

    // 라벨 감지
    labels, err := client.DetectLabels(ctx, image, nil, 10)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to detect label",
		})
        return
    }

    var filteredLabelsInfo []map[string]interface{}

    // 필터링 및 추가
    for _, label := range labels {
        if _, found := filterKeywords[label.Description]; found {
            labelInfo := map[string]interface{}{
                "description": label.Description,
                "score":       label.Score,
            }
            filteredLabelsInfo = append(filteredLabelsInfo, labelInfo)
        }
    }

    // 신뢰도 점수에 따라 내림차순으로 정렬
    sort.Slice(filteredLabelsInfo, func(i, j int) bool {
        return filteredLabelsInfo[i]["score"].(float32) > filteredLabelsInfo[j]["score"].(float32)
    })

    // 감지된 라벨 출력
    c.JSON(http.StatusOK, gin.H{
		"message": "Upload successful",
        "labels":filteredLabelsInfo,
	})
}
