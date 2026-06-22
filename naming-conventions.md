  Case = visibility. This is the one rule Go actually enforces.
  - ExportedName (capital) — public, visible outside the package
  - unexportedName (lowercase) — package-private
  - Applies to every identifier: funcs, types, vars, constants, struct fields, methods.
  
  MixedCaps, not snake_case or SCREAMING_SNAKE.
  - parseToken, HTTPClient, MaxRetries — yes
  - parse_token, MAX_RETRIES — no, even for constants

  Acronyms stay all-caps.
  - HTTPServer, parseURL, userID, JSONDecoder
  - not HttpServer, parseUrl, userId
  
  Package names: short, lowercase, single word, no underscores.
  - drive, instagram, queue, httpd — yes
  - instagram_api, instagramAPI, InstagramAPI — no
  - Avoid stuttering — if the package is queue, name the type queue.Queue only if there's no better
   word; prefer queue.FS or queue.Manager.
   
  Files: lowercase, underscores allowed.
  - token_refresh.go, carousel.go — yes
  - Test files must end _test.go. 
  
  Receivers: short, consistent, usually one or two letters.
  - func (q *Queue) Push(...) not func (queue *Queue) Push(...)
  - Same receiver name across every method on the type.
  
  Interfaces: -er suffix when it's a single-method interface.
  - Reader, Publisher, Refresher
  - Bigger interfaces just get descriptive names: Storage, Client.

  Errors:
  - Sentinel error variables: ErrNotFound, ErrInvalidToken (prefix Err)
  - Error types: *ParseError (suffix Error)

  Getters drop the Get prefix.
  - user.Name() not user.GetName()
  - Setters keep Set: user.SetName(...)
  
  Constants follow the same rules as vars — MaxRetries, not MAX_RETRIES.

  Don't repeat the package name in identifiers.
  - instagram.Client, not instagram.InstagramClient
  - queue.New(), not queue.NewQueue()
  
  gofmt enforces formatting automatically (run on save), and go vet + staticcheck will yell about
  the naming conventions it can detect. Lean on the tooling — it'll teach you the style faster than
   memorizing rules.