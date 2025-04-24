import fs from 'fs';

// Bubble Sort function
function bubbleSort(arr) {
  let i = 0;
  while (i < arr.length - 1) {
    if (arr[i] > arr[i + 1]) {
      [arr[i], arr[i + 1]] = [arr[i + 1], arr[i]];
      i = 0; // Restart sorting after swap
    } else {
      i++;
    }
  }
  return arr;
}

// Lambda handler function
export const handler = async (event) => {
  const inputFiles = [
    './pi_input_25.json',
    './pi_input_50.json',
    './pi_input_100.json',
    './pi_input_500.json',
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
      const inputArray = jsonData.numbers || []; // Extract the array field

      console.log(`Input data for ${jsonFilePath}:`, inputArray); // Debug log for input data

      // Process the input array
      const sortedArray = bubbleSort(inputArray); // Sort the array

      // Store the result for this file
      results[jsonFilePath] = sortedArray;

      // Log the sorted array for this file
      console.log(`Sorted array for ${jsonFilePath}:`, sortedArray);
    } catch (error) {
      console.error(`Error processing file ${inputFiles[i]}:`, error.message);
    }
  }

  // Return results for all files, including sorted arrays
  return {
    body: JSON.stringify({
      sortedArrays: results,
    }),
  };
};
