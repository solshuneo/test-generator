curl -X POST http://localhost:3001/number \
   -H "Content-Type: application/yaml" \
   --data-binary $'min: 10\nmax: 100'
