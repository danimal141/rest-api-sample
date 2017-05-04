const gulp = require('gulp')
const aglio = require('aglio')
const gaglio = require('gulp-aglio')
const rename = require('gulp-rename')
const fs = require('fs')
const includePath = process.cwd() + '/api'
const paths = aglio.collectPathsSync(fs.readFileSync('api/all.apib', {encoding: 'utf8'}), includePath)

gulp.task('build', () =>
  gulp.src('api/all.apib')
    .pipe(gaglio({template: 'default'}))
    .pipe(rename('out.html'))
    .pipe(gulp.dest('docs'))
)

gulp.task('watch', () =>
  gulp.watch(paths, ['build'])
)

gulp.task('default', ['build', 'watch'])
