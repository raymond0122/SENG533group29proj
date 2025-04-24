import fs from 'fs';

function letterCountNode(inputString) {
  const letterCounts = new Array(26).fill(0);

  for (let i = 0; i < inputString.length; i++) {
    const charCode = inputString[i].charCodeAt(0);
    if (charCode >= 97 && charCode <= 122) {
      letterCounts[charCode - 97]++;
    }
  }

  return letterCounts;
}

// Lambda handler function
export const handler = async (event) => {
  const inputFiles = [
    './string_input_25.json',
    './string_input_50.json',
    './string_input_100.json',
    './string_input_500.json',
  ];

  // Initialize an object to store results for all files
  const results = {};

  // Iterate over the input files
  for (let i = 0; i < inputFiles.length; i++) {
    try {
      const jsonFilePath = inputFiles[i];

      // Read and parse the JSON file
      const data = fs.readFileSync(jsonFilePath, 'utf8');
      const jsonData = JSON.parse(data);
      const inputString = jsonData.data || ''; // Extract the text field

      // Process the input string
      const result = letterCountNode(inputString.toLowerCase()); // Convert to lowercase

      // Store the result for this file
      results[jsonFilePath] = result;

      console.log(`Letter counts for ${jsonFilePath}:`, result);
    } catch (error) {
      console.error(`Error processing file ${inputFiles[i]}:`, error.message);
    }
  }

  // Return results for all files
  return {
    body: JSON.stringify({
      counts: results,
    }),
  };
};
