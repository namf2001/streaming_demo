console.time('Streaming fetch time');
fetch('http://localhost:8080/all_records')
    .then(response => {
            const reader = response.body.getReader();
            const decoder = new TextDecoder();
            let buffer = '';

            function read() {
                reader.read().then(({done, value}) => {
                    if (done) {
                        return;
                    }
                    buffer += decoder.decode(value, {stream: true});

                    while (true) {
                        const start = buffer.indexOf('\n');
                        if (start == -1) break;
                        const line = buffer.substring(0, start);
                        buffer = buffer.substring(start + 1);
                        try {
                            const record = JSON.parse(line);
                            console.log(record);
                        } catch (e) {
                            console.error('Error parsing:', e);
                        }
                    }

                    read();
                });
            }

            read();
            console.timeEnd('Non-streaming fetch time');
        }
    )
    .catch(error => {
        console.error('Error:', error);
    });