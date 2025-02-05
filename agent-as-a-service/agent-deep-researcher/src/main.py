'''
Created on Feb 5, 2025

@author: autonomous
'''

from flask import Flask, jsonify, request 
from flask_restful import Resource, Api 
from multiprocessing import Process
from deep_research import run_deep_research

app = Flask(__name__) 
api = Api(app) 
  
class Hello(Resource): 
    def get(self): 
        return jsonify({'message': 'hello world'}) 
  
class DeepResearch(Resource): 
    def get(self): 
        query = request.args.get('query')
        run_id = request.args.get('id')
        try:
            p = Process(target=run_deep_research, args=(run_id, query, 1))
            p.start()
            return jsonify({'message': 'ok'}) 
        except Exception as ex:
            return jsonify({'message': 'failed'})
         
        
api.add_resource(Hello, '/') 
api.add_resource(DeepResearch, '/research') 
  
  
if __name__ == '__main__': 
    app.run(debug = True) 