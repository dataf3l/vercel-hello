package handler

import (
	"fmt"
	"net/http"
	"time"
)

// Handler is the main entry point for the Vercel serverless function
func Handler(w http.ResponseWriter, r *http.Request) {
	// Set content type to HTML
	w.Header().Set("Content-Type", "text/html")

	// Get current time
	currentTime := time.Now().Format("2006-01-02 15:04:05 MST")

	// Create a simple HTML response
	html := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <title>Go Hello World on Vercel</title>
    <style>
        body { 
            font-family: Arial, sans-serif; 
            text-align: center; 
            margin-top: 100px;
            background: linear-gradient(135deg, #667eea 0%%, #764ba2 100%%);
            color: white;
        }
        .container {
            background: rgba(255,255,255,0.1);
            padding: 40px;
            border-radius: 10px;
            display: inline-block;
            backdrop-filter: blur(10px);
        }
        h1 { color: #fff; margin-bottom: 20px; }
        p { font-size: 18px; margin: 10px 0; }
        .time { font-style: italic; color: #f0f0f0; }
        .nav-links { margin-top: 30px; }
        .nav-links a { 
            color: #fff; 
            text-decoration: none; 
            margin: 0 15px; 
            padding: 10px 20px;
            background: rgba(255,255,255,0.2);
            border-radius: 5px;
            display: inline-block;
            margin-bottom: 10px;
        }
        .nav-links a:hover { background: rgba(255,255,255,0.3); }
    </style>
</head>
<body>
    <div class="container">
        <h1>🚀 Hello World from Go on Vercel!</h1>
        <p>This is a minimalistic Go serverless function with multiple endpoints!</p>
        <p>Method: <strong>%s</strong></p>
        <p>Path: <strong>%s</strong></p>
        <p class="time">Server time: %s</p>
        <p>✨ Deployed successfully! ✨</p>
        
        <div class="nav-links">
            <h2 style="color: #fff; margin-bottom: 20px;">🔗 Available Routes:</h2>
            <a href="/api/items">📦 Items API</a>
            <a href="/api/about">📋 About</a>
            <a href="/api/demo">🎮 Demo</a>
        </div>
    </div>
</body>
</html>`, r.Method, r.URL.Path, currentTime)

	// Write the response
	fmt.Fprint(w, html)
}
