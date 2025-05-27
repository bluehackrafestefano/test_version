1. Clone the repo

2. Delete .git folder

3. Push initial to the repository

4. Understand code structure:
- Directory structure:
```
├───backend
│   ├───config
│   ├───controllers
│   ├───data
│   │   ├───images
│   │   ├───invoice
│   │   └───util
│   ├───middlewares
│   │   ├───common
│   │   ├───helpers
│   │   ├───user_actions
│   │   └───validator
│   ├───models
│   ├───public
│   │   ├───css
│   │   └───js
│   ├───routes
│   └───utils
├───build
│   └───static
│       ├───css
│       ├───js
│       └───media
├───public
└───src
    ├───components
│   └───static
│       ├───css
│       ├───js
│   └───static
│   └───static
│       ├───css
│       ├───js
│       └───media
├───public
└───src
    ├───components
    │   ├───misc
    │   ├───navigations
    │   └───pages
    ├───datas
    │   └───faqs
    ├───images
    │   └───New folder
    └───layout
```

- This is a Node.js full-stack web application, likely using Express.js for the backend and React.js for the frontend.
```txt
├───backend        # Node.js / Express backend
├───build          # Production build of the frontend (compiled React app)
├───public         # Public assets (probably for React)
└───src            # Source code for the React frontend
```

### backend/ — The Express.js Backend
Contains all the server-side logic.
- config/
Configuration files like database connection, environment variables, or third-party APIs.

- controllers/
Logic for handling HTTP requests. Each file typically maps to a route (e.g., authController.js, invoiceController.js).

- data/
Raw or semi-structured data used by the backend.
  - images/: Media assets.
  - invoice/: Likely invoice templates or related data files.
  - util/: Utility data or scripts (e.g., constants, test datasets).

- middlewares/
Custom middleware functions.
  - common/: General middleware (e.g., logging, error handling).
  - helpers/: Utility functions used in middleware.
  - user_actions/: Auth or user-specific middlewares.
  - validator/: Middleware for validating inputs (e.g., using Joi or express-validator).

- models/
Database models (e.g., MongoDB schemas if using Mongoose, or Sequelize models).

- public/
Static files served by Express (like legacy HTML/CSS/JS assets).
  - css/, js/: Client-side files, likely not part of React.

- routes/
API route definitions (e.g., userRoutes.js, invoiceRoutes.js), linking URLs to controller logic.

- utils/
Helper functions, utilities (e.g., date formatting, token generation, file handling).

### build/ — Production Build of Frontend
This is the output from npm run build for React. It's static files ready to be served.

- static/css, static/js, static/media: Minified styles, JS bundles, and media files.

### public/ — React Public Assets
Used during development and build of React app. Anything here is accessible at the root URL (e.g., favicon.ico, index.html).

### src/ — React Frontend Source Code
All the development code for the frontend app.

- components/
Reusable UI components.
  - misc/: Miscellaneous components.
  - navigations/: Navbar, sidebar, breadcrumbs, etc.
  - pages/: Page-specific components or views.

- datas/faqs/
Static FAQ content or JSON used in the app.

- images/
Image assets for the React app.

- layout/
Common layout components like Header, Footer, or main App wrapper.

### package.json
The package.json file is the metadata and configuration file for a Node.js project. It is essential for both backend (Node.js/Express) and frontend (React) projects that use Node.js.
- Defines the project (name, version, description)
- Lists dependencies (libraries the project needs)
- Specifies scripts (commands you run with npm run)
- Describes the environment (Node version, module format, etc.)

5. Set up the environment (Docker, virtualenv, etc.)
### Needs
- node 22
- python 3.10
- docker
- npm

### Create virtualenv
- Create a `.dockerignore` file to exclude unnecessary files from the Docker build context.

- Create a virtualenv
```py
python -m venv env # for Windows or
py -3.10 -m venv env # for Windows
virtualenv env # for Mac/Linux or;
virtualenv env -p python3 # for Mac/Linux
```

- Add env/ to the dockerignore and gitignore files

- Activate scripts:
```bash
.\env\Scripts\activate  # for Windows
source env/bin/activate  # for MAC/Linux
```
- Run the app locally:
```sh
npm install
```

### Create a Dockerfile
- Use docker linting locally; hadolint
- Build using docker
```sh
docker build -t stefanorafe/app:v1 .
docker login
docker push stefanorafe/app:v1
docker run --name my_app -d -p 3000:3000 stefanorafe/app:v1
```
