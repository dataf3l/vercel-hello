# Go Hello World API on Vercel

A minimalistic Go API deployed on Vercel with multiple endpoints and interactive features.

## 🚀 Live Deployment

**Working URL:** https://aa-eight-xi.vercel.app  
**Latest Deployment:** https://api-no25pr7n4-dataf3ls-projects.vercel.app (requires auth)

## 📋 Available Routes

### ✅ 1. **/** (Root) - Redirects to Hello World
- **URL:** https://aa-eight-xi.vercel.app/
- **Status:** ✅ WORKING
- **Description:** Main landing page with navigation to other routes

### ✅ 2. **/api/hello** - Hello World Page
- **URL:** https://aa-eight-xi.vercel.app/api/hello
- **Status:** ✅ WORKING
- **Description:** Interactive hello world page with route navigation
- **Features:** Server time display, responsive design

### ❌ 3. **/api/items** - JSON API
- **URL:** https://aa-eight-xi.vercel.app/api/items
- **Status:** ❌ NOT DEPLOYED (404 Error)
- **Description:** Returns JSON array of 5 hardcoded items
- **Content-Type:** application/json
- **CORS:** Enabled for cross-origin requests

### ❌ 4. **/api/about** - About Page
- **URL:** https://aa-eight-xi.vercel.app/api/about
- **Status:** ❌ NOT DEPLOYED (404 Error)
- **Description:** Information about the API and technical stack
- **Features:** Technical details, endpoint documentation

### ❌ 5. **/api/demo** - Interactive Demo
- **URL:** https://aa-eight-xi.vercel.app/api/demo
- **Status:** ❌ NOT DEPLOYED (404 Error)
- **Description:** Demo page with JavaScript fetch functionality
- **Features:** Fetches items from /api/items using JavaScript, interactive UI

## 🛠️ Technical Stack

- **Language:** Go 1.19
- **Platform:** Vercel Serverless Functions
- **Framework:** Native Go net/http
- **Deployment:** Vercel CLI

## 🧪 Testing the Routes

Use the provided test script to verify working routes:

```bash
chmod +x test_working_routes.sh
./test_working_routes.sh
```

Or test individual working routes with curl:

```bash
# Test main page (WORKING)
curl https://aa-eight-xi.vercel.app/

# Test hello page (WORKING)
curl https://aa-eight-xi.vercel.app/api/hello

# Test JSON API (NOT DEPLOYED - 404)
curl https://aa-eight-xi.vercel.app/api/items

# Test about page (NOT DEPLOYED - 404)
curl https://aa-eight-xi.vercel.app/api/about

# Test demo page (NOT DEPLOYED - 404)
curl https://aa-eight-xi.vercel.app/api/demo
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