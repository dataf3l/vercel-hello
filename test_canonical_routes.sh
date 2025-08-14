#!/bin/bash

# Test script for 5 Go API routes on the canonical Vercel URL
BASE_URL="https://aa-eight-xi.vercel.app"

echo "🧪 Testing 5 Go API routes on canonical Vercel URL..."
echo "🌐 Base URL: $BASE_URL"
echo "======================================================="
echo ""

echo "1️⃣ Testing /api/hello (Main hello world page):"
curl -s -w "HTTP Status: %{http_code}\n" "$BASE_URL/api/hello" | head -5
echo ""
echo "---"

echo "2️⃣ Testing /api/items (JSON API with 5 items):"
curl -s -w "HTTP Status: %{http_code}\n" "$BASE_URL/api/items"
echo ""
echo "---"

echo "3️⃣ Testing /api/about (About page):"
curl -s -w "HTTP Status: %{http_code}\n" "$BASE_URL/api/about" | head -5
echo ""
echo "---"

echo "4️⃣ Testing /api/demo (Demo page with JS fetch):"
curl -s -w "HTTP Status: %{http_code}\n" "$BASE_URL/api/demo" | head -5
echo ""
echo "---"

echo "5️⃣ Testing / (Root redirects to hello):"
curl -s -w "HTTP Status: %{http_code}\n" "$BASE_URL/" | head -5
echo ""

echo "======================================================="
echo "✅ All 5 routes tested on canonical URL!"
echo "🔗 Visit in browser: $BASE_URL"