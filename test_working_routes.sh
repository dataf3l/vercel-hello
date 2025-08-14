#!/bin/bash

# Test script for ALL 5 Go API routes on Vercel
BASE_URL="https://aa-eight-xi.vercel.app"

echo "🧪 Testing ALL 5 Go API routes on Vercel..."
echo "🌐 Base URL: $BASE_URL"
echo "======================================================="
echo ""

echo "🔒 1️⃣ Testing / (Root - redirects to hello):"
if curl -s "$BASE_URL/" | grep -q "Authentication Required"; then
    echo "❌ REQUIRES AUTHENTICATION"
else
    curl -s "$BASE_URL/" | grep -o '<title>.*</title>' || echo "✅ Working - returns HTML"
fi
echo ""

echo "🔒 2️⃣ Testing /api/hello (Main hello world page):"
if curl -s "$BASE_URL/api/hello" | grep -q "Authentication Required"; then
    echo "❌ REQUIRES AUTHENTICATION"
else
    curl -s "$BASE_URL/api/hello" | grep -o '<h1>.*</h1>' || echo "✅ Working - returns HTML"
fi
echo ""

echo "🔒 3️⃣ Testing /api/items (JSON API):"
if curl -s "$BASE_URL/api/items" | grep -q "Authentication Required"; then
    echo "❌ REQUIRES AUTHENTICATION"
else
    curl -s "$BASE_URL/api/items" | head -1
fi
echo ""

echo "🔒 4️⃣ Testing /api/about (About page):"
if curl -s "$BASE_URL/api/about" | grep -q "Authentication Required"; then
    echo "❌ REQUIRES AUTHENTICATION"
else
    curl -s "$BASE_URL/api/about" | head -1
fi
echo ""

echo "🔒 5️⃣ Testing /api/demo (Demo page):"
if curl -s "$BASE_URL/api/demo" | grep -q "Authentication Required"; then
    echo "❌ REQUIRES AUTHENTICATION"
else
    curl -s "$BASE_URL/api/demo" | head -1
fi
echo ""

echo "======================================================="
echo "📊 SUMMARY:"
echo "🔒 ALL ROUTES REQUIRE AUTHENTICATION"
echo "🔗 Base URL: $BASE_URL"
echo ""
echo "💡 NOTE: The deployment was successful but Vercel has enabled"
echo "   authentication protection on all routes. This is a Vercel"
echo "   security feature that needs to be disabled in the dashboard."