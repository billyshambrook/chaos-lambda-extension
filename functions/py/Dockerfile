
FROM extension-chaos as chaos-extension

FROM public.ecr.aws/lambda/python:3.8

COPY --from=chaos-extension /opt/extensions/ /opt/extensions/

COPY app.py requirements.txt ./

RUN python3 -m pip install -r requirements.txt -t .

# Command can be overwritten by providing a different command in the template directly.
CMD ["app.lambda_handler"]
