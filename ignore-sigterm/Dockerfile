FROM fedora
RUN echo 'while true; do date; echo "Hello world!"; sleep 10; done' > script.sh && chmod +x script.sh
ENTRYPOINT ["/bin/bash", "/script.sh"]
