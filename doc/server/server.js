// required libraries
const express = require('express');
const helmet = require('helmet');
const rateLimit = require('express-rate-limit');
const cors = require('cors');
const { body, validationResult } = require('express-validator');
const morgan = require('morgan');
const dotenv = require('dotenv');
const fs = require('fs');
const path = require('path');

// Load variables from .env file
dotenv.config();

// Initialize Express
const app = express();

// Use Helmet to set security-related HTTP headers
app.use(helmet());

// Enable CORS for specific origin
const corsOptions = {
    origin: process.env.ALLOWED_ORIGIN,
};
app.use(cors(corsOptions));

// rate limiting to prevent small DDoS attacks
const limiter = rateLimit({
    windowMs: 15 * 60 * 1000, // 15 minutes
    max: 99, // Limit each IP to 99 requests
});
app.use(limiter);

// Middleware to parse JSON requests
app.use(express.json());

// Setup logging implementing Morgan
app.use(morgan('combined'));

// Sample route with input validation
app.post('/submit', [
    body('username').isAlphanumeric().withMessage('Username must be alphanumeric'),
    body('email').isEmail().withMessage('Invalid email format'),
], (req, res) => {
    // Validate request data
    const errors = validationResult(req);
    if (!errors.isEmpty()) {
        return res.status(400).json({ errors: errors.array() });
    }

    // Here you would handle the request data, e.g., save to database
    res.status(200).send('Data received successfully');
});

// Error handling middleware
app.use((err, req, res, next) => {
    console.error(err.stack); // Log the error details
    res.status(500).send('Internal Server Error'); // Send a generic error message
});

// Handle 404 errors
app.use((req, res) => {
    res.status(404).send('Not Found');
});

// Start server
const PORT = process.env.PORT || 3000;
app.listen(PORT, () => {
    console.log(`Server is running on http://localhost:${PORT}`);
});
