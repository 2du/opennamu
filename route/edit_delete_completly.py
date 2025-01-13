from .tool.func import *

def edit_delete_completly(name = 'Test'):
    with get_db_connect() as conn:
        curs = conn.cursor()

        if acl_check('', 'owner_auth', '', '') == 1:
            return re_error(conn, 3)

        if flask.request.method == 'POST':
            acl_check(tool = 'owner_auth', memo = 'completly delete ' + name)

            curs.execute(db_change("delete from history where title = ?"), [name])
            curs.execute(db_change("delete from acl where title = ?"), [name])
            curs.execute(db_change("delete from back where title = ?"), [name])
            curs.execute(db_change("delete from back where link = ?"), [name])
            curs.execute(db_change("delete from data where title = ?"), [name])
            curs.execute(db_change("delete from data_set where doc_name = ?"), [name])
            curs.execute(db_change("delete from rc where title = ?"), [name])
            curs.execute(db_change("delete from history where title = ?"), [name])
            curs.execute(db_change("delete from user_set where name = 'star_doc' and data = ?"), [name])
            curs.execute(db_change("delete from rd where title = ?"), [name])

            return redirect(conn, '/w/' + url_pas(name))
        else:
            return easy_minify(conn, flask.render_template(skin_check(conn),
                imp = [name, wiki_set(conn), wiki_custom(conn), wiki_css(['(' + get_lang(conn, 'delete_completly') + ')', 0])],
                data = '''
                    <form method="post">
                        <span>''' + get_lang(conn, 'delete_warning') + '''</span>
                        <hr class="main_hr">
                        <button type="submit">''' + get_lang(conn, 'delete') + '''</button>
                    </form>
                ''',
                menu = [['w/' + url_pas(name), get_lang(conn, 'return')]]
            ))
