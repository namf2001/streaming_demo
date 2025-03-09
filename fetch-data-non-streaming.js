console.time('Non-streaming fetch time');
fetch('http://localhost:8080/all_records_non_streaming')
    .then(response => response.json())
    .then(data => {
        console.log('Received all records');
        console.log(data);
        console.timeEnd('Non-streaming fetch time');
        // Process data as needed
    })
    .catch(error => {
        console.error('Error:', error);
        console.timeEnd('Non-streaming fetch time');
    });