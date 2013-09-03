module.exports = function(grunt) {

	// Project configuration.
	grunt.initConfig({

		pkg: grunt.file.readJSON('package.json'),

		jshint: {
			all: ['Gruntfile.js', 'src/js/app/**/*.js', 'test/**/*.js']
		},

		concat: {
			options: {
				separator: ';',
			},
			dist: {
				src: ['src/js/lib/**/*.js', 'src/js/app/**/*.js'],
				dest: 'public/js/app.js',
			},
		},

		uglify: {
			options: {
				banner: '/*! <%= pkg.name %> <%= grunt.template.today("yyyy-mm-dd") %> */\n'
			},
			build: {
				src: 'src/<%= pkg.name %>.js',
				dest: 'public/<%= pkg.name %>.min.js'
			}
		},

		copy: {
			main: {
				files: [{
					expand: true,
					flatten: true,
					src: ['src/index.html'],
					dest: 'public/',
					// filter: 'isFile'
				}, // includes files in path
				]
			}
		},

		watch: {
			files: ['main.go'],
			tasks: ['reload']
		}

	});

	// Load the plugin that provides the "uglify" task.
	grunt.loadNpmTasks('grunt-contrib-jshint');
	grunt.loadNpmTasks('grunt-contrib-uglify');
	grunt.loadNpmTasks('grunt-contrib-concat');
	grunt.loadNpmTasks('grunt-contrib-copy');
	grunt.loadNpmTasks('grunt-reload');

	grunt.registerTask('default', ['jshint', 'concat', 'copy']);

};