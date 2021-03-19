from aiocqhttp import Event
from aiocqhttp.default import on_message, run


@on_message
async def handle_msg(event: Event):
    print(event.message)


if __name__ == '__main__':
    run(host='127.0.0.1', port=8081)
