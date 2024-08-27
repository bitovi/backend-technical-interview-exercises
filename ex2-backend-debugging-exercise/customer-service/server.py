from flask import Flask, jsonify
import psycopg2
import os

app = Flask(__name__)

def run_query(query):
	try:
		connection_string = os.getenv('DATABASE_URL')
		conn = psycopg2.connect(connection_string)
		cursor = conn.cursor()

		cursor.execute(query)
		columns = [desc[0] for desc in cursor.description]
		results = cursor.fetchall()

		conn.close()

		formatted_results = [dict(zip(columns, row)) for row in results]
		return formatted_results
	except Exception as e:
		return str(e)

@app.route("/customers", methods=["GET"])
def execute_query():
	try:
		data = run_query("select * from get_customer_details()")
		return jsonify(data), 200
	except Exception as e:
		return jsonify({"error": str(e)}), 500

if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0', port=3001)