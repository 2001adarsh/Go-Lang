# Go-Lang
Go Crash Course

### References to learn from:

**Youtube**: https://www.youtube.com/watch?v=YS4e4q9oBaU

**Blogs**: https://www.digitalocean.com/community/tutorial_series/how-to-code-in-go 

-----

### Structuring of GO Applications:

1. Layered
  - handlers
    - bears.go
    - reviews.go
     
  - models
    - bear.go
    - review.go
    - storage.go
    
  - storage
    - json.go
    - memory.go
    
  - main.go
    - data.go


2. Grouping by Modules:
    - bears
      - bear.go
      - handler.go
    - database
      - data.go
      - json.go
      - memory.go
    - reviews
      - handler.go
      - review.go
    - main.go
  
3. Domain Driven Development
    - Establish your domain and buisness logic
    - Define your bound context(s), the models within each context and ubiquitous language
    - Categorise building block of your system.
      - Entity
      - value object
      - domain event
      - Aggregation
      - service
      - repository
      - factory
4. Hexagonal Architecture
    - Core domain surrounded by interfaces to external inputs and outputs
    - Dependencies pointing only inwards
      - cmd
        - data
          - main.go
        - server
          - main.go
      - pkg
        - adding
        - bears
        - listing
        - reviewing
        - reviews
      - storage
        - json
        - memory.go
        - types.go
      - gopkg.toml
   
   -------------
   
   - Keep _test.go files next to main files.
   - Use shared mock sub package
  
