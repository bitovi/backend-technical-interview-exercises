# Bitovi Express Example

## Scenario

The Sales department head requests the ability to fetch their team's sales data using a customizable filter that excludes sales data below the amount they pass in.

Your task is to implement the filter.

Detailed instructions are located in the `index.js` file, but at a high level:

Step 1 is to set up the parameter in the `/sales` handler located in `index.js` with proper validation.

Step 2 is to use that parameter to filter the data.

## Command line operations

The interviewer will start the server.
Note that the server will automatically restart after changes have been performed.

At any time, call the server with a command like this:
```
curl 'localhost:3002/sales'
```

## Interviewer setup instructions

```
npm start
```