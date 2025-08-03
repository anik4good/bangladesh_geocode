$baseUrl = "https://bdgeo.root2tech.com"

while ($true) {
    # Test all endpoints
    Write-Host "Testing /api/divisions"
    curl -X GET "$baseUrl/api/divisions"
    
    Write-Host "Testing /api/division/1"
    curl -X GET "$baseUrl/api/division/1"
    
    Write-Host "Testing /api/districts"
    curl -X GET "$baseUrl/api/districts"
    
    Write-Host "Testing /api/division/rangpur"
    curl -X GET "$baseUrl/api/division/rangpur"
    
    Write-Host "Testing /api/division/rangpur/panchagarh"
    curl -X GET "$baseUrl/api/division/rangpur/panchagarh"
    
    Write-Host "Testing /api/division/rangpur/panchagarh/debiganj"
    curl -X GET "$baseUrl/api/division/rangpur/panchagarh/debiganj"
    
    # Write-Host "Testing /metrics"
    # curl -X GET "$baseUrl/metrics"
    
    # Wait 5 seconds before next iteration
    Start-Sleep -Seconds 1
}