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