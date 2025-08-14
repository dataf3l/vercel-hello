package handler

import (
	"fmt"
	"net/http"
	"time"
)

// Handler provides information about the API
func Handler(w http.ResponseWriter, r *http.Request) {
	// Set content type to HTML
	w.Header().Set("Content-Type", "text/html")

	// Get current time
	currentTime := time.Now().Format("2006-01-02 15:04:05 MST")

	// Create an about page
	html := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <title>About - Go API on Vercel</title>
    <style>
        body { 
            font-family: Arial, sans-serif; 
            text-align: center; 
            margin: 20px;
            background: linear-gradient(135deg, #74b9ff 0%%, #0984e3 100%%);
            color: white;
            min-height: 100vh;
        }
        .container {
            background: rgba(255,255,255,0.1);
            padding: 40px;
            border-radius: 15px;
            display: inline-block;
            backdrop-filter: blur(10px);
            margin-top: 50px;
            max-width: 600px;
        }
        h1 { color: #fff; margin-bottom: 30px; }
        h2 { color: #f0f0f0; margin-top: 30px; }
        p { font-size: 16px; margin: 15px 0; line-height: 1.6; }
        .time { font-style: italic; color: #ddd; }
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
        <h1>📋 About This API</h1>
        <p>This is a minimalistic Go API deployed on Vercel with multiple endpoints!</p>
        
        <h2>🚀 Available Endpoints:</h2>
        <p><strong>/api/hello</strong> - Main hello world page</p>
        <p><strong>/api/items</strong> - JSON API with 5 sample items</p>
        <p><strong>/api/about</strong> - This about page</p>
        
        <h2>🛠️ Technical Stack:</h2>
        <p>• <strong>Language:</strong> Go 1.19</p>
        <p>• <strong>Platform:</strong> Vercel Serverless Functions</p>
        <p>• <strong>Framework:</strong> Native Go net/http</p>
        
        <p class="time">Server time: %s</p>
        
        <div class="nav-links">
            <a href="/api/hello">🏠 Home</a>
            <a href="/api/items">📦 Items API</a>
            <a href="/api/demo">🎮 Demo Page</a>
        </div>
    </div>
</body>
</html>`, currentTime)

	// Write the response
	fmt.Fprint(w, html)
}
