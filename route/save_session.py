from .tool.func import *

def save_session():
    data = flask.request.get_json()
    redirect_url = data.get('redirectUrl', '/')

    path = '/' + redirect_url.split('://')[-1].split('/', 1)[-1]
    # Flask의 서버 세션에 저장
    flask.session['redirectUrl'] = path

    return {'message': 'Session value saved'}