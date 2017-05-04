

	var statusLabel = $("#status");
	function sendUpSignal() {
		$.ajax( {
			url: "/up",
			type: "POST",
			success: function(data) {
				console.log("Up request sent");
				statusLabel.html("<font color='green'>Moving</font>");
			},
			error: function(data) {
				console.log(data);
				statusLabel.html("<font color='green'>Moving</font>");
			}
		});
	}
	function sendDownSignal() {
		$.ajax( {
			url: "/down",
			type: "POST",
			success: function(data) {
				console.log("Down request sent");
				statusLabel.html("<font color='green'>Moving</font>");
			}
		});
	}
	function sendLeftSignal() {
		$.ajax( {
			url: "/left",
			type: "POST",
			success: function(data) {
				console.log("Left request sent");
				statusLabel.html("<font color='green'>Moving</font>");
			}
		});
	}
	function sendRightSignal() {
		$.ajax( {
			url: "/right",
			type: "POST",
			success: function(data) {
				console.log("Right request sent");
				statusLabel.html("<font color='green'>Moving</font>");
			}
		});
	}

	function sendStopSignal() {
		$.ajax( {
			url: "/stop",
			type: "POST",
			success: function(data) {
				console.log("Stop request sent");
				statusLabel.html("<font color='red'>Stopped</font>");

			}
		})
	}
