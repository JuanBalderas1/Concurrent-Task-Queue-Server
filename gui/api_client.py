import json
from urllib import error, request


class APIError(Exception):
    pass


class ChimeraAPIClient:
    def __init__(self, base_url="http://localhost:8080"):
        self.base_url = base_url.rstrip("/")

    def get_members(self):
        return self._send_request(
            method="GET",
            path="/members",
        )

    def create_member(self, name, contact, role):
        payload = {
            "name": name,
            "contact": contact,
            "role": role,
        }

        return self._send_request(
            method="POST",
            path="/members",
            payload=payload,
        )

    def get_tasks(self):
        return self._send_request(
            method="GET",
            path="/tasks",
        )

    def create_task(
        self,
        sender_id,
        recipient_id,
        task_type,
        payload,
        max_retries=3,
    ):
        task_data = {
            "sender_id": sender_id,
            "recipient_id": recipient_id,
            "type": task_type,
            "payload": payload,
            "max_retries": max_retries,
        }

        return self._send_request(
            method="POST",
            path="/tasks",
            payload=task_data,
        )

    def _send_request(
        self,
        method,
        path,
        payload=None,
    ):
        url = f"{self.base_url}{path}"
        body = None

        headers = {
            "Accept": "application/json",
        }

        if payload is not None:
            body = json.dumps(payload).encode("utf-8")
            headers["Content-Type"] = "application/json"

        api_request = request.Request(
            url=url,
            data=body,
            headers=headers,
            method=method,
        )

        try:
            with request.urlopen(
                api_request,
                timeout=2,
            ) as response:
                response_body = response.read()

                if not response_body:
                    return None

                return json.loads(
                    response_body.decode("utf-8")
                )

        except error.HTTPError as http_error:
            message = http_error.read().decode(
                "utf-8"
            ).strip()

            if not message:
                message = (
                    f"Server returned HTTP "
                    f"{http_error.code}"
                )

            raise APIError(message) from http_error

        except error.URLError as connection_error:
            raise APIError(
                "Unable to connect to Chimera Task Server."
            ) from connection_error
        