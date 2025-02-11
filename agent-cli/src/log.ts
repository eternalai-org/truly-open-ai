// import chalk from 'chalk';
import chalk from 'chalk';
import Table from 'cli-table3';

function getTimestamp() {
    return new Date().toLocaleTimeString();
}

function logInfo(message: string): void {
    console.log(chalk.blue(`[${getTimestamp()}] [INFO] ${message}`));
}

function logSuccess(message: string): void {
    console.log(chalk.green(`[${getTimestamp()}] [SUCCESS] ${message}`));
}

function logError(message: string): void {
    console.log(chalk.red(`[${getTimestamp()}] [ERROR] ${message}`));
}

function logWarning(message: string): void {
    console.log(chalk.yellow(`[${getTimestamp()}] [WARNING] ${message}`));
}

// Function to log headers (e.g., for sections) in bold cyan
function logHeader(message: string): void {
    console.log(chalk.cyan.bold(`[${getTimestamp()}] [HEADER] ${message}`));
}

// Function to log a title with a line break for better readability
function logTitle(message: string): void {
    console.log(chalk.bold.white(message));
    console.log(chalk.white('-'.repeat(message.length)));  // Creates a separator line
}

// interface Agent {
//     AgentID: string
//     Name: string
//     NetworkName: string
//     ChainID: string
//     ModelID: string
// }

// interface Column {
//     header: string
//     width: number
// }

// const HeaderInfoMap: Record<string, Column> = {
//     "AgentID": {
//         header: "Agent ID",
//         width: 10
//     },
//     "Name": {
//         header: "Name",
//         width: 20
//     },

//     "AgentID": {
//         header: "Agent ID",
//         width: 10
//     },
//     "AgentID": {
//         header: "Agent ID",
//         width: 10
//     },
//     "AgentID": {
//         header: "Agent ID",
//         width: 10
//     },

// }

// Function to format and display a table
function logTable(data: Array<Record<string, any>>): void {
    if (data.length === 0) {
        console.log(chalk.yellow('No data to display.'));
        return;
    }

    // Get the table headers from the keys of the first object
    const headers = Object.keys(data[0]);

    // Calculate the width for each column (max length of header or data)
    const colWidths = headers.map(header =>
        Math.max(...data.map(item => item[header].toString().length), header.length)
    );

    console.log("colWidths: ", colWidths);

    // Create the table's horizontal border line
    const borderLine = `+${colWidths.map(width => '-'.repeat(width + 2)).join('+')}+`;

    // Create a function to format each row
    const formatRow = (row: any) => {
        return `| ${headers.map((header, index) => {
            const value = row[header].toString();
            const padding = ' '.repeat(colWidths[index] - value.length);
            return value + padding;
        }).join(' | ')} |`;
    };

    // Function to format the header row with color
    const formatHeaderRow = (headers: string[]) => {
        return `| ${headers.map((header, index) => {
            const value = chalk.cyan.bold(header.toUpperCase()); // Use blue and bold for header
            const padding = ' '.repeat(colWidths[index] - header.length);
            // const padding = "";
            return value + padding;
        }).join(' | ')} |`;
    };

    // Create a new table instance with the headers
    const table = new Table({
        head: headers.map(header => chalk.bold(header)),  // Bold headers
        colWidths: colWidths,  // Set a default width for columns
        style: {
            head: ['bold', 'cyan'],  // Style for the headers
            border: ['grey'],  // Style for the borders
        }
    });

    // Add each row of data to the table
    data.forEach(row => {
        const rowData = headers.map(header => row[header] ?? 'N/A');
        table.push(rowData);
    });

    // Print the formatted table
    // console.log(table.toString());
    console.log(borderLine);
    console.log(formatHeaderRow(headers)); // Print header
    console.log(borderLine);
    data.forEach(row => console.log(formatRow(row)));
    console.log(borderLine);


}


export {
    logInfo,
    logSuccess,
    logError,
    logWarning,
    logHeader,
    logTitle,
    logTable,


}
