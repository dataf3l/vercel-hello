package handler

import (
	"fmt"
	"net/http"
	"time"
)

// Handler provides a demo page with JavaScript to fetch items
func Handler(w http.ResponseWriter, r *http.Request) {
	// Set content type to HTML
	w.Header().Set("Content-Type", "text/html")
	
	// Get current time
	currentTime := time.Now().Format("2006-01-02 15:04:05 MST")
	
	// Create a demo page with JavaScript
	html := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <title>Demo - Go API on Vercel</title>
    <style>
        body { 
            font-family: Arial, sans-serif; 
            margin: 20px;
            background: linear-gradient(135deg, #a29bfe 0%%, #6c5ce7 100%%);
            color: white;
            min-height: 100vh;
        }
        .container {
            background: rgba(255,255,255,0.1);
            padding: 40px;
            border-radius: 15px;
            backdrop-filter: blur(10px);
            margin: 20px auto;
            max-width: 800px;
        }
        h1 { color: #fff; margin-bottom: 30px; text-align: center; }
        h2 { color: #f0f0f0; margin-top: 30px; }
        .items-grid {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
            gap: 20px;
            margin-top: 20px;
        }
        .item-card {
            background: rgba(255,255,255,0.15);
            padding: 20px;
            border-radius: 10px;
            border: 1px solid rgba(255,255,255,0.2);
        }
        .item-name { font-weight: bold; font-size: 18px; margin-bottom: 10px; }
        .item-price { color: #fdcb6e; font-weight: bold; font-size: 16px; }
        .item-category { color: #ddd; font-size: 14px; }
        .item-description { margin: 10px 0; line-height: 1.4; }
        .loading { text-align: center; font-style: italic; }
        .error { color: #ff7675; text-align: center; }
        .nav-links { text-align: center; margin-top: 30px; }
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
        .fetch-button {
            background: rgba(255,255,255,0.2);
            color: white;
            border: none;
            padding: 12px 24px;
            border-radius: 8px;
            cursor: pointer;
            font-size: 16px;
            margin: 20px 0;
            display: block;
            margin-left: auto;
            margin-right: auto;
        }
        .fetch-button:hover { background: rgba(255,255,255,0.3); }
        .time { text-align: center; font-style: italic; color: #ddd; margin-top: 20px; }
    </style>
</head>
<body>
    <div class="container">
        <h1>🎮 Demo Page - Fetch Items with JavaScript</h1>
        
        <p style="text-align: center;">This page demonstrates fetching JSON data from our Go API using JavaScript fetch().</p>
        
        <button class="fetch-button" onclick="fetchItems()">🔄 Fetch Items from API</button>
        
        <div id="items-container">
            <div class="loading">Click the button above to fetch items from /api/items</div>
        </div>
        
        <h2>📡 API Endpoints Available:</h2>
        <ul style="line-height: 2;">
            <li><strong>/api/hello</strong> - Main hello world page</li>
            <li><strong>/api/items</strong> - JSON API with 5 sample items</li>
            <li><strong>/api/about</strong> - About page with information</li>
            <li><strong>/api/demo</strong> - This demo page</li>
        </ul>
        
        <div class="nav-links">
            <a href="/api/hello">🏠 Home</a>
            <a href="/api/items">📦 Items JSON</a>
            <a href="/api/about">📋 About</a>
        </div>
        
        <p class="time">Server time: %s</p>
    </div>

    <script>
        async function fetchItems() {
            const container = document.getElementById('items-container');
            container.innerHTML = '<div class="loading">Loading items...</div>';
            
            try {
                const response = await fetch('/api/items');
                
                if (!response.ok) {
                    throw new Error('Network response was not ok: ' + response.status);
                }
                
                const items = await response.json();
                
                // Create HTML for items
                let html = '<h2>📦 Fetched Items:</h2><div class="items-grid">';
                
                items.forEach(item => {
                    html += ` + "`" + `
                        <div class="item-card">
                            <div class="item-name">${item.name}</div>
                            <div class="item-description">${item.description}</div>
                            <div class="item-category">Category: ${item.category}</div>
                            <div class="item-price">$${item.price}</div>
                        </div>
                    ` + "`" + `;
                });
                
                html += '</div>';
                container.innerHTML = html;
                
            } catch (error) {
                container.innerHTML = '<div class="error">❌ Error fetching items: ' + error.message + '</div>';
                console.error('Error:', error);
            }
        }
        
        // Optional: Auto-fetch on page load
        // fetchItems();
    </script>
</body>
</html>`, currentTime)

	// Write the response
	fmt.Fprint(w, html)
}