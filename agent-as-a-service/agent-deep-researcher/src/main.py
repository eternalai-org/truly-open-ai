from flask import Flask, jsonify, request
from flask_restful import Resource, Api
from multiprocessing import Process
from deep_research import run_deep_research

app = Flask(__name__)
api = Api(app)

class Hello(Resource):
    def get(self):
        return jsonify({'result': 'hello world'})

class DeepResearch(Resource):
    def post(self):
        req_data = request.get_json()
        query = req_data.get('query')
        req_id = req_data.get('req_id')
        try:
            p = Process(target=run_deep_research, args=(req_id, query, 1))
            p.start()
            return jsonify({'result': 'ok'})
        except Exception as ex:
            return jsonify({'result': 'failed'})


api.add_resource(Hello, '/')
api.add_resource(DeepResearch, '/research')


if __name__ == '__main__':
    app.run(debug = True)
