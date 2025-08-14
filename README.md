# Go Hello World API on Vercel

A minimalistic Go API deployed on Vercel with multiple endpoints and interactive features.

## 🚀 Live Deployment

**Canonical URL:** https://api-dataf3ls-projects.vercel.app  
**Current Working URL:** https://api-8917k3gaw-dataf3ls-projects.vercel.app

## 📋 Available Routes

### 1. **/** (Root) - Redirects to Hello World
- **URL:** https://api-dataf3ls-projects.vercel.app/
- **Description:** Main landing page with navigation to other routes

### 2. **/api/hello** - Hello World Page
- **URL:** https://api-dataf3ls-projects.vercel.app/api/hello
- **Description:** Interactive hello world page with route navigation
- **Features:** Server time display, responsive design

### 3. **/api/items** - JSON API
- **URL:** https://api-dataf3ls-projects.vercel.app/api/items
- **Description:** Returns JSON array of 5 hardcoded items
- **Content-Type:** application/json
- **CORS:** Enabled for cross-origin requests

### 4. **/api/about** - About Page
- **URL:** https://api-dataf3ls-projects.vercel.app/api/about
- **Description:** Information about the API and technical stack
- **Features:** Technical details, endpoint documentation

### 5. **/api/demo** - Interactive Demo
- **URL:** https://api-dataf3ls-projects.vercel.app/api/demo
- **Description:** Demo page with JavaScript fetch functionality
- **Features:** Fetches items from /api/items using JavaScript, interactive UI

## 🛠️ Technical Stack

- **Language:** Go 1.19
- **Platform:** Vercel Serverless Functions
- **Framework:** Native Go net/http
- **Deployment:** Vercel CLI

## 🧪 Testing the Routes

Use the provided test script to verify all routes:

```bash
chmod +x tmp_rovodev_test_routes.sh
./tmp_rovodev_test_routes.sh
```

Or test individual routes with curl:

```bash
# Test JSON API
curl https://api-dataf3ls-projects.vercel.app/api/items

# Test main page
curl https://api-dataf3ls-projects.vercel.app/

# Test about page
curl https://api-dataf3ls-projects.vercel.app/api/about

# Test demo page
curl https://api-dataf3ls-projects.vercel.app/api/demo

# Test hello page
curl https://api-dataf3ls-projects.vercel.app/api/hello
```

## 📦 Sample JSON Response

The `/api/items` endpoint returns:

```json
[
  {
    "id": 1,
    "name": "Wireless Headphones",
    "description": "High-quality wireless headphones with noise cancellation",
    "category": "Electronics",
    "price": 199.99
  },
  {
    "id": 2,
    "name": "Coffee Mug",
    "description": "Ceramic coffee mug with heat retention technology",
    "category": "Kitchen",
    "price": 24.99
  },
  // ... 3 more items
]
```

## 🎨 Features

- **Responsive Design:** All pages are mobile-friendly
- **Interactive Demo:** JavaScript fetch demonstration
- **CORS Support:** API endpoints support cross-origin requests
- **Error Handling:** Proper HTTP status codes and error responses
- **Real-time Data:** Server timestamps on all pages

## 🚀 Deployment

The project is automatically deployed to Vercel using:

```bash
npx vercel --prod --yes
```

## 📁 Project Structure

```
.
├── api/
│   ├── hello.go    # Main hello world handler
│   ├── items.go    # JSON API handler
│   ├── about.go    # About page handler
│   └── demo.go     # Interactive demo handler
├── vercel.json     # Vercel configuration
├── package.json    # Node.js dependencies
└── README.md       # This file
```