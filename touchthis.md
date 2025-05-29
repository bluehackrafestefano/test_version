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
- go version shows at least go1.20.x.

- Create .env file
- 
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


### Terraform

- Create terraform/
terraform init
- Create main.tf file
- aws provider and ec2 instance add
terraform fmt
terraform validate
terraform plan -out plan.txt
terraform apply
terraform destroy
- add test/ to main dir
- add ec2_test.go
- initialize go in proj dir
```sh
go mod init test_version
go get github.com/gruntwork-io/terratest/modules/terraform
go get github.com/gruntwork-io/terratest/modules/aws
go get github.com/aws/aws-sdk-go/service/ec2
go get github.com/stretchr/testify/assert
go test ./test -v
go test ./test -v -count=1
```

### AWS CLI

### s3
aws s3 ls
aws s3api create-bucket --bucket my-task-bucket-12355
aws s3 cp ./test2.txt s3://my-task-bucket-12355
aws s3 sync . s3://my-task-bucket-12355 --exclude "env/*" --exclude "node_modules/*"
aws s3api list-objects --bucket my-task-bucket-12355
aws s3 rm --recursive s3://my-task-bucket-12355
aws s3api delete-bucket --bucket my-task-bucket-12355

### ec2
aws ec2 run-instances --image-id ami-0953476d60561c955 --count 1 --instance-type t2.micro --key-name rafe --security-group-ids sg-072cefb3b4a4b6732 --profile cw --region us-east-1 --tag-specifications 'ResourceType=instance,Tags=[{Key=Name,Value=MyInstance}]'

aws ec2 describe-instances --filters "Name=instance-state-name,Values=running" --profile cw --region us-east-1 --output table

aws ec2 describe-instances --filters "Name=instance-state-name,Values=running" --profile cw --region us-east-1 --output table | grep InstanceId

aws ec2 terminate-instances --instance-ids i-085c41f55d634aec7 --profile cw --region us-east-1

### lambda
aws lambda list-functions --profile cw --region us-east-1

index.js
```
exports.handler = async (event) => {
    console.log("Event:", event);
    return { statusCode: 200, body: "Hello from Lambda!" };
};
```
manually zip index.js
trust-policy.json
```
{
  "Version": "2012-10-17",
  "Statement": [{
    "Effect": "Allow",
    "Principal": { "Service": "lambda.amazonaws.com" },
    "Action": "sts:AssumeRole"
  }]
}
```
aws iam create-role \
  --role-name lambda-execution-role \
  --assume-role-policy-document file://trust-policy.json

aws iam attach-role-policy \
  --role-name lambda-execution-role \
  --policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole

aws iam list-roles

aws iam list-roles
arn:aws:iam::121099113991:role/lambda-execution-role

aws lambda create-function \
  --function-name my-lambda-function \
  --runtime nodejs18.x \
  --role arn:aws:iam::121099113991:role/lambda-execution-role \
  --handler index.handler \
  --zip-file fileb://index.zip

aws lambda get-function --function-name my-lambda-function

aws lambda invoke \
  --function-name my-lambda-function \
  --payload '{}' \
  response.json

cat response.json

### Logging

Fluentd
npm install winston
npm start 

- Create logger.js
```js
const { createLogger, format, transports } = require('winston');

const logger = createLogger({
  level: 'debug',
  format: format.combine(
    format.timestamp(),
    format.simple()
  ),
  transports: [
    new transports.File({ filename: 'app.log' }),
    new transports.Console()
  ],
});

module.exports = logger;
```

- that will create app.log file under proj dir

backend/app.js
```js
app.use((req, res, next) => {
  logger.info(`Incoming request: ${req.method} ${req.originalUrl}`);
  next();
});

app.use((req, res, next) => {
  logger.warn(`404 Not Found: ${req.method} ${req.originalUrl}`);
  res.status(404).json({ error: 'Not Found' });
});
```

- replace in backend/server.js
```js
    console.log(`Error: ${err.message}`);
```
- to:
```js
    logger.info(`Server running on port ${PORT}`);
```


### Prometheus node exporter
1. Install prom-client
npm install prom-client
2. Add a /metrics endpoint to the backend/app.js
```js
const client = require('prom-client');
const register = new client.Registry();
client.collectDefaultMetrics({ register });

// Expose /metrics endpoint
app.get('/metrics', async (req, res) => {
  res.set('Content-Type', register.contentType);
  res.end(await register.metrics());
});

// Optional: custom metric example
const httpRequestCounter = new client.Counter({
  name: 'http_requests_total',
  help: 'Total number of HTTP requests',
  labelNames: ['method', 'route', 'status'],
});
register.registerMetric(httpRequestCounter);

// Middleware to increment the counter
app.use((req, res, next) => {
  res.on('finish', () => {
    httpRequestCounter.inc({ method: req.method, route: req.originalUrl, status: res.statusCode });
  });
  next();
});
```
- You can get your Prometheus metrics by visiting:
http://localhost:3099/metrics

Example output:
```
# HELP process_cpu_user_seconds_total Total user CPU time spent in seconds.
# TYPE process_cpu_user_seconds_total counter
process_cpu_user_seconds_total 29.063

# HELP process_cpu_system_seconds_total Total system CPU time spent in seconds.
# TYPE process_cpu_system_seconds_total counter
process_cpu_system_seconds_total 0.953

# HELP process_cpu_seconds_total Total user and system CPU time spent in seconds.
# TYPE process_cpu_seconds_total counter
process_cpu_seconds_total 30.016

# HELP process_start_time_seconds Start time of the process since unix epoch in seconds.
# TYPE process_start_time_seconds gauge
process_start_time_seconds 1748529858

# HELP process_resident_memory_bytes Resident memory size in bytes.
# TYPE process_resident_memory_bytes gauge
process_resident_memory_bytes 115613696

# HELP nodejs_eventloop_lag_seconds Lag of event loop in seconds.
# TYPE nodejs_eventloop_lag_seconds gauge
nodejs_eventloop_lag_seconds 0

# HELP nodejs_eventloop_lag_min_seconds The minimum recorded event loop delay.
# TYPE nodejs_eventloop_lag_min_seconds gauge
nodejs_eventloop_lag_min_seconds 0.00991232

# HELP nodejs_eventloop_lag_max_seconds The maximum recorded event loop delay.
# TYPE nodejs_eventloop_lag_max_seconds gauge
nodejs_eventloop_lag_max_seconds 3.061841919

# HELP nodejs_eventloop_lag_mean_seconds The mean of the recorded event loop delays.
# TYPE nodejs_eventloop_lag_mean_seconds gauge
nodejs_eventloop_lag_mean_seconds 0.08344710868292683

# HELP nodejs_eventloop_lag_stddev_seconds The standard deviation of the recorded event loop delays.
# TYPE nodejs_eventloop_lag_stddev_seconds gauge
nodejs_eventloop_lag_stddev_seconds 0.15226345537488747

# HELP nodejs_eventloop_lag_p50_seconds The 50th percentile of the recorded event loop delays.
# TYPE nodejs_eventloop_lag_p50_seconds gauge
nodejs_eventloop_lag_p50_seconds 0.072024063

# HELP nodejs_eventloop_lag_p90_seconds The 90th percentile of the recorded event loop delays.
# TYPE nodejs_eventloop_lag_p90_seconds gauge
nodejs_eventloop_lag_p90_seconds 0.104595455

# HELP nodejs_eventloop_lag_p99_seconds The 99th percentile of the recorded event loop delays.
# TYPE nodejs_eventloop_lag_p99_seconds gauge
nodejs_eventloop_lag_p99_seconds 0.164888575

# HELP nodejs_active_resources Number of active resources that are currently keeping the event loop alive, grouped by async resource type.
# TYPE nodejs_active_resources gauge
nodejs_active_resources{type="FSReqCallback"} 1
nodejs_active_resources{type="PipeWrap"} 5
nodejs_active_resources{type="TCPServerWrap"} 1
nodejs_active_resources{type="ProcessWrap"} 3
nodejs_active_resources{type="TCPSocketWrap"} 1
nodejs_active_resources{type="Timeout"} 4
nodejs_active_resources{type="Immediate"} 2

# HELP nodejs_active_resources_total Total number of active resources.
# TYPE nodejs_active_resources_total gauge
nodejs_active_resources_total 17

# HELP nodejs_active_handles Number of active libuv handles grouped by handle type. Every handle type is C++ class name.
# TYPE nodejs_active_handles gauge
nodejs_active_handles{type="Socket"} 6
nodejs_active_handles{type="Server"} 1
nodejs_active_handles{type="ChildProcess"} 3

# HELP nodejs_active_handles_total Total number of active handles.
# TYPE nodejs_active_handles_total gauge
nodejs_active_handles_total 10

# HELP nodejs_active_requests Number of active libuv requests grouped by request type. Every request type is C++ class name.
# TYPE nodejs_active_requests gauge
nodejs_active_requests{type="FSReqCallback"} 1

# HELP nodejs_active_requests_total Total number of active requests.
# TYPE nodejs_active_requests_total gauge
nodejs_active_requests_total 1

# HELP nodejs_heap_size_total_bytes Process heap size from Node.js in bytes.
# TYPE nodejs_heap_size_total_bytes gauge
nodejs_heap_size_total_bytes 69873664

# HELP nodejs_heap_size_used_bytes Process heap size used from Node.js in bytes.
# TYPE nodejs_heap_size_used_bytes gauge
nodejs_heap_size_used_bytes 44872024

# HELP nodejs_external_memory_bytes Node.js external memory size in bytes.
# TYPE nodejs_external_memory_bytes gauge
nodejs_external_memory_bytes 22374833

# HELP nodejs_heap_space_size_total_bytes Process heap space size total from Node.js in bytes.
# TYPE nodejs_heap_space_size_total_bytes gauge
nodejs_heap_space_size_total_bytes{space="read_only"} 0
nodejs_heap_space_size_total_bytes{space="new"} 33554432
nodejs_heap_space_size_total_bytes{space="old"} 28078080
nodejs_heap_space_size_total_bytes{space="code"} 1572864
nodejs_heap_space_size_total_bytes{space="shared"} 0
nodejs_heap_space_size_total_bytes{space="trusted"} 3170304
nodejs_heap_space_size_total_bytes{space="new_large_object"} 0
nodejs_heap_space_size_total_bytes{space="large_object"} 3342336
nodejs_heap_space_size_total_bytes{space="code_large_object"} 155648
nodejs_heap_space_size_total_bytes{space="shared_large_object"} 0
nodejs_heap_space_size_total_bytes{space="trusted_large_object"} 0

# HELP nodejs_heap_space_size_used_bytes Process heap space size used from Node.js in bytes.
# TYPE nodejs_heap_space_size_used_bytes gauge
nodejs_heap_space_size_used_bytes{space="read_only"} 0
nodejs_heap_space_size_used_bytes{space="new"} 10021752
nodejs_heap_space_size_used_bytes{space="old"} 27212264
nodejs_heap_space_size_used_bytes{space="code"} 1361408
nodejs_heap_space_size_used_bytes{space="shared"} 0
nodejs_heap_space_size_used_bytes{space="trusted"} 2891928
nodejs_heap_space_size_used_bytes{space="new_large_object"} 0
nodejs_heap_space_size_used_bytes{space="large_object"} 3253408
nodejs_heap_space_size_used_bytes{space="code_large_object"} 138304
nodejs_heap_space_size_used_bytes{space="shared_large_object"} 0
nodejs_heap_space_size_used_bytes{space="trusted_large_object"} 0

# HELP nodejs_heap_space_size_available_bytes Process heap space size available from Node.js in bytes.
# TYPE nodejs_heap_space_size_available_bytes gauge
nodejs_heap_space_size_available_bytes{space="read_only"} 0
nodejs_heap_space_size_available_bytes{space="new"} 6472840
nodejs_heap_space_size_available_bytes{space="old"} 320072
nodejs_heap_space_size_available_bytes{space="code"} 112960
nodejs_heap_space_size_available_bytes{space="shared"} 0
nodejs_heap_space_size_available_bytes{space="trusted"} 220664
nodejs_heap_space_size_available_bytes{space="new_large_object"} 16777216
nodejs_heap_space_size_available_bytes{space="large_object"} 0
nodejs_heap_space_size_available_bytes{space="code_large_object"} 0
nodejs_heap_space_size_available_bytes{space="shared_large_object"} 0
nodejs_heap_space_size_available_bytes{space="trusted_large_object"} 0

# HELP nodejs_version_info Node.js version info.
# TYPE nodejs_version_info gauge
nodejs_version_info{version="v22.16.0",major="22",minor="16",patch="0"} 1

# HELP nodejs_gc_duration_seconds Garbage collection duration by kind, one of major, minor, incremental or weakcb.
# TYPE nodejs_gc_duration_seconds histogram
nodejs_gc_duration_seconds_bucket{le="0.001",kind="incremental"} 1
nodejs_gc_duration_seconds_bucket{le="0.01",kind="incremental"} 1
nodejs_gc_duration_seconds_bucket{le="0.1",kind="incremental"} 1
nodejs_gc_duration_seconds_bucket{le="1",kind="incremental"} 1
nodejs_gc_duration_seconds_bucket{le="2",kind="incremental"} 1
nodejs_gc_duration_seconds_bucket{le="5",kind="incremental"} 1
nodejs_gc_duration_seconds_bucket{le="+Inf",kind="incremental"} 1
nodejs_gc_duration_seconds_sum{kind="incremental"} 0.0009644000008702278
nodejs_gc_duration_seconds_count{kind="incremental"} 1
nodejs_gc_duration_seconds_bucket{le="0.001",kind="major"} 0
nodejs_gc_duration_seconds_bucket{le="0.01",kind="major"} 1
nodejs_gc_duration_seconds_bucket{le="0.1",kind="major"} 1
nodejs_gc_duration_seconds_bucket{le="1",kind="major"} 1
nodejs_gc_duration_seconds_bucket{le="2",kind="major"} 1
nodejs_gc_duration_seconds_bucket{le="5",kind="major"} 1
nodejs_gc_duration_seconds_bucket{le="+Inf",kind="major"} 1
nodejs_gc_duration_seconds_sum{kind="major"} 0.0062836000025272365
nodejs_gc_duration_seconds_count{kind="major"} 1
nodejs_gc_duration_seconds_bucket{le="0.001",kind="minor"} 107
nodejs_gc_duration_seconds_bucket{le="0.01",kind="minor"} 263
nodejs_gc_duration_seconds_bucket{le="0.1",kind="minor"} 281
nodejs_gc_duration_seconds_bucket{le="1",kind="minor"} 281
nodejs_gc_duration_seconds_bucket{le="2",kind="minor"} 281
nodejs_gc_duration_seconds_bucket{le="5",kind="minor"} 281
nodejs_gc_duration_seconds_bucket{le="+Inf",kind="minor"} 281
nodejs_gc_duration_seconds_sum{kind="minor"} 0.7629821000173689
nodejs_gc_duration_seconds_count{kind="minor"} 281

# HELP http_requests_total Total number of HTTP requests
# TYPE http_requests_total counter
```

- ![monitor on grafana](image.png)
https://grafana.com/grafana/dashboards/1860-node-exporter-full/
