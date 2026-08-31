curl -X POST http://localhost:3001/number \
   -H "Content-Type: application/yaml" \
   --data-binary $'min: 10\nmax: 100'
curl -X POST https://test-generator.shuneo.com/number \
    -H "Content-Type: application/yaml" \
    --data-binary $'min: 10\nmax: 100'
