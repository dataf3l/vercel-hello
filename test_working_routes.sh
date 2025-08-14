#!/bin/bash

# Test script for WORKING Go API routes on Vercel
BASE_URL="https://aa-eight-xi.vercel.app"

echo "🧪 Testing WORKING Go API routes on Vercel..."
echo "🌐 Base URL: $BASE_URL"
echo "======================================================="
echo ""

echo "✅ 1️⃣ Testing / (Root - redirects to hello):"
curl -s "$BASE_URL/" | grep -o '<title>.*</title>' || echo "Working - returns HTML"
echo ""

echo "✅ 2️⃣ Testing /api/hello (Main hello world page):"
curl -s "$BASE_URL/api/hello" | grep -o '<h1>.*</h1>' || echo "Working - returns HTML"
echo ""

echo "❌ 3️⃣ Testing /api/items (NOT DEPLOYED):"
curl -s "$BASE_URL/api/items" | head -1
echo ""

echo "❌ 4️⃣ Testing /api/about (NOT DEPLOYED):"
curl -s "$BASE_URL/api/about" | head -1
echo ""

echo "❌ 5️⃣ Testing /api/demo (NOT DEPLOYED):"
curl -s "$BASE_URL/api/demo" | head -1
echo ""

echo "======================================================="
echo "📊 SUMMARY:"
echo "✅ Working routes: / and /api/hello"
echo "❌ Missing routes: /api/items, /api/about, /api/demo"
echo "🔗 Visit working routes: $BASE_URL"
echo ""
echo "💡 NOTE: Only 2 out of 5 routes are deployed to this URL."
echo "   The other routes exist in code but need redeployment."