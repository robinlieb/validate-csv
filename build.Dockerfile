FROM scratch
COPY validate-csv /usr/local/bin/validate-csv
COPY run.sh /
RUN chmod +x /run.sh
ENTRYPOINT [ "./run.sh" ]