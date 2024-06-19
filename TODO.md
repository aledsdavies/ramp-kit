# TODO

## Initial Setup

- [x] Initialize chi for routing
- [x] Basic Templ setup to render with Go
- [x] Connect CSS and serve static files
- [ ] Set up global error handling

## Serving Static Files

- [x] Introduce gzip for static content
    - [ ] Investigate switching out gzip for brotli
        - [ ] Test Brotli compression effectiveness on various file types
        - [ ] Ensure server compatibility and configuration for Brotli
- [x] Introduce hashing into the `/styles/{hash}/*` path to allow us to break caching on file changes
    - [x] Implement hash generation for static files
    - [x] Update references to static files to use hashed paths
    - [X] Investigate having a unique hash per file instead of one for all css
- [ ] Introduce images, etc., to the project, ideally using modern web standards
    - [ ] Optimize images for the web
    - [ ] Implement responsive images (e.g., using `srcset` and `sizes`)
    - [ ] Consider using modern formats like WebP and AVIF

## Styling

- [x] Investigate a nice way to load css per page
    - [X] Investigate CSS Modules
    - [x] Investigate pulling CSS from Templ definitions
- [x] Embed styles into their own global var and serve them from `/styles/*`
- [X] Customize a reset.css using modern standards and practices
    - [X] Research modern CSS resets (e.g., Normalize.css, modern-css-reset)
    - [X] Implement and integrate a customized reset.css
- [ ] Manage reading post CSS modules into the templ templates
    - [X] Mapping the generated CSS class names into templ components
    - [ ] Dynamic loading of the css. If a value from a css module is loaded we should
    also load it in the HTML in the template
    - [ ] Panic Error if we try to load a CSS value that does not exist

- [ ] Create a main.css for the project
    - [ ] Produce a design system using CSS vars to control the project from main.css
        - [ ] Structure a type system with CSS vars
        - [ ] Define color palette using OKLCH for supporting P3 colors
        - [x] Set up typography (font families, sizes, weights)
        - [ ] Establish spacing scale (margins, paddings)
        - [ ] Create reusable components (buttons, forms, cards)

## Database

- [ ] Build a module for connecting to databases
- [ ] Add support out of the box for some kind of SQLite solution
- [ ] Add support for in memory datastores that we can use for local testing

## Permissions

- [ ] Build a module for managing permissions in the app
- [ ] Possible embed spicedb to take advantage of using a zanzibar type permissions models
    - [ ]

## Pages

- [ ] Investigate the benefits of encapsulating Templ functionality within a struct to serve as the page builder
    - **Modularity**: Assess how wrapping Templ into a struct can improve code modularity, allowing for better separation of concerns and easier maintenance.
    - **Reusability**: Evaluate the potential for reusing the struct-based page builder across different pages, reducing duplication and simplifying updates.
    - **Initialization**: Examine the benefits of having a single point of initialization for page-related data and dependencies, streamlining the setup process.
    - **Customization**: Determine how the struct can facilitate easy customization and extension of pages, providing a flexible foundation for various page types.
    - **Integration**: Explore how integrating additional features (e.g., middleware, common components) within the struct can enhance the overall functionality and consistency of pages.
    - **Performance**: Analyze any performance improvements that may result from encapsulating Templ logic, such as more efficient rendering and data handling.
- [ ] Investigate how we should support components in the project
- [ ] Write Templ HTML for the login page with HTMX and Alpine.js
    - [ ] Design login form
    - [ ] Implement interactive elements with Alpine.js
    - [ ] Use HTMX for dynamic form handling
- [ ] Write Templ HTML for the landing page once logged in
    - [ ] Design dashboard layout
    - [ ] Integrate data visualization (charts, tables)
    - [ ] Implement interactive components with Alpine.js and HTMX


## Setting Up Authentication

- [ ] Setup module so that we can inject different authentication options into the project
- [ ] Set up basic authentication by trying different options
    - [ ] Research authentication libraries and methods (JWT, OAuth, sessions)
    - [ ] Implement authentication backend
    - [ ] Create middleware for protected routes
    - [ ] Build login and registration forms
    - [ ] Test authentication flow and error handling

## Additional Features

- [ ] Implement user roles and permissions
    - [ ] Define user roles (admin, user, guest)
    - [ ] Set up role-based access control (RBAC) with zanzibar
    - [ ] Test permissions and access restrictions
- [ ] Set up database integration
    - [ ] Modularize the DB to allow switching to different providers
    - [ ] Choose a database (PostgreSQL, MySQL, SQLite)
    - [ ] Design database schema
    - [ ] Implement CRUD operations
    - [ ] Integrate the database with the authentication system

