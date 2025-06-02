# PowerShell script to test the Notes API

# Base URL
$baseUrl = "http://localhost:3001"

# Function to display formatted JSON
function Format-Json {
    param([Parameter(Mandatory, ValueFromPipeline)][String] $json)
    $indent = 0
    ($json -Replace '(?=(?:(?:[^"]*"){2})*[^"]*$)[\[\{]', "`n`${0}`n$(' ' * ($indent += 2))" -Replace '(?=(?:(?:[^"]*"){2})*[^"]*$)[\]\}]', "`n$(' ' * ($indent -= 2))`${0}" -Replace '(?=(?:(?:[^"]*"){2})*[^"]*$),', "`${0}`n$(' ' * $indent)")
}

# Function to make HTTP requests
function Invoke-ApiRequest {
    param(
        [string]$method,
        [string]$endpoint,
        [hashtable]$headers = @{},
        [object]$body = $null,
        [string]$token = ""
    )

    $uri = "$baseUrl$endpoint"
    
    # Add content type if body is provided
    if ($body) {
        $headers["Content-Type"] = "application/json"
    }
    
    # Add authorization header if token is provided
    if ($token) {
        $headers["Authorization"] = "Bearer $token"
    }
    
    $params = @{
        Method = $method
        Uri = $uri
        Headers = $headers
    }
    
    # Add body if provided
    if ($body) {
        $jsonBody = $body | ConvertTo-Json
        $params["Body"] = $jsonBody
    }
    
    try {
        $response = Invoke-RestMethod @params
        return $response
    } catch {
        $statusCode = $_.Exception.Response.StatusCode.value__
        $errorDetails = $_.ErrorDetails.Message
        
        Write-Host "Error: $statusCode" -ForegroundColor Red
        if ($errorDetails) {
            Write-Host $errorDetails -ForegroundColor Red
        } else {
            Write-Host $_.Exception.Message -ForegroundColor Red
        }
        return $null
    }
}

# Test server connection
Write-Host "`n🔍 Testing server connection..." -ForegroundColor Cyan
$root = Invoke-ApiRequest -method "GET" -endpoint "/"
if ($root) {
    Write-Host "✅ Server is running: $root" -ForegroundColor Green
} else {
    Write-Host "❌ Server not running. Please start the server with 'go run main.go'" -ForegroundColor Red
    exit
}

# Register a new user
Write-Host "`n🔍 Registering a new user..." -ForegroundColor Cyan
$registerBody = @{
    name = "Test User"
    email = "test@example.com"
    password = "password123"
}
$registerResponse = Invoke-ApiRequest -method "POST" -endpoint "/register" -body $registerBody
if ($registerResponse) {
    Write-Host "✅ User registered successfully:" -ForegroundColor Green
    $registerResponse | ConvertTo-Json | Format-Json | Write-Host
} else {
    Write-Host "⚠️ User might already exist, trying to login..." -ForegroundColor Yellow
}

# Login
Write-Host "`n🔍 Logging in..." -ForegroundColor Cyan
$loginBody = @{
    email = "test@example.com"
    password = "password123"
}
$loginResponse = Invoke-ApiRequest -method "POST" -endpoint "/login" -body $loginBody
if ($loginResponse) {
    $token = $loginResponse.token
    Write-Host "✅ Login successful, received token:" -ForegroundColor Green
    Write-Host $token -ForegroundColor Green
} else {
    Write-Host "❌ Login failed. Exiting." -ForegroundColor Red
    exit
}

# Create a note
Write-Host "`n🔍 Creating a new note..." -ForegroundColor Cyan
$createNoteBody = @{
    title = "Test Note"
    content = "This is a test note created by the API test script."
}
$createNoteResponse = Invoke-ApiRequest -method "POST" -endpoint "/notes" -body $createNoteBody -token $token
if ($createNoteResponse) {
    $noteId = $createNoteResponse.ID
    Write-Host "✅ Note created successfully:" -ForegroundColor Green
    $createNoteResponse | ConvertTo-Json | Format-Json | Write-Host
} else {
    Write-Host "❌ Failed to create note. Exiting." -ForegroundColor Red
    exit
}

# Get all notes
Write-Host "`n🔍 Fetching all notes..." -ForegroundColor Cyan
$getNotesResponse = Invoke-ApiRequest -method "GET" -endpoint "/notes" -token $token
if ($getNotesResponse) {
    Write-Host "✅ Notes retrieved successfully:" -ForegroundColor Green
    $getNotesResponse | ConvertTo-Json -Depth 4 | Format-Json | Write-Host
} else {
    Write-Host "❌ Failed to retrieve notes." -ForegroundColor Red
}

# Get a single note
Write-Host "`n🔍 Fetching a single note..." -ForegroundColor Cyan
$getNoteResponse = Invoke-ApiRequest -method "GET" -endpoint "/notes/$noteId" -token $token
if ($getNoteResponse) {
    Write-Host "✅ Note retrieved successfully:" -ForegroundColor Green
    $getNoteResponse | ConvertTo-Json | Format-Json | Write-Host
} else {
    Write-Host "❌ Failed to retrieve note." -ForegroundColor Red
}

# Update a note
Write-Host "`n🔍 Updating a note..." -ForegroundColor Cyan
$updateNoteBody = @{
    title = "Updated Test Note"
    content = "This note has been updated by the API test script."
}
$updateNoteResponse = Invoke-ApiRequest -method "PUT" -endpoint "/notes/$noteId" -body $updateNoteBody -token $token
if ($updateNoteResponse) {
    Write-Host "✅ Note updated successfully:" -ForegroundColor Green
    $updateNoteResponse | ConvertTo-Json | Format-Json | Write-Host
} else {
    Write-Host "❌ Failed to update note." -ForegroundColor Red
}

# Delete a note
Write-Host "`n🔍 Deleting a note..." -ForegroundColor Cyan
$deleteNoteResponse = Invoke-ApiRequest -method "DELETE" -endpoint "/notes/$noteId" -token $token
if ($deleteNoteResponse) {
    Write-Host "✅ Note deleted successfully:" -ForegroundColor Green
    $deleteNoteResponse | ConvertTo-Json | Format-Json | Write-Host
} else {
    Write-Host "❌ Failed to delete note." -ForegroundColor Red
}

Write-Host "`n✅ API test completed successfully!" -ForegroundColor Green
