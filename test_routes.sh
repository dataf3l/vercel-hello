#!/bin/bash

# Test script for 5 Go API routes on Vercel
BASE_URL="https://api-8917k3gaw-dataf3ls-projects.vercel.app"

echo "Testing 5 Go API routes on Vercel..."
echo "Base URL: $BASE_URL"
echo ""

echo "1. Testing /api/hello (Main hello world page):"
curl -s "$BASE_URL/api/hello"
echo ""

echo "2. Testing /api/items (JSON API with 5 items):"
curl -s "$BASE_URL/api/items"
echo ""

echo "3. Testing /api/about (About page):"
curl -s "$BASE_URL/api/about"
echo ""

echo "4. Testing /api/demo (Demo page with JS fetch):"
curl -s "$BASE_URL/api/demo"
echo ""

echo "5. Testing / (Root redirects to hello):"
curl -s "$BASE_URL/"
echo ""

echo "All 5 routes tested!"