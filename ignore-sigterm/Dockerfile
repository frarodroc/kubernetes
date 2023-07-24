FROM fedora
RUN echo 'date; echo "Hello world!"; sleep 10' > script.sh && chmod +x script.sh
ENTRYPOINT ["/bin/bash", "/script.sh"]
